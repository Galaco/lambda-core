package bsp

import (
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/dispinfo"
	"github.com/galaco/bsp/primitives/dispvert"
	"github.com/galaco/bsp/primitives/face"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/bsp/primitives/visibility"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/filesystem"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/galaco/go-me-engine/valve/libwrapper/stringtable"
	"github.com/galaco/go-me-engine/valve/libwrapper/vpk"
	material2 "github.com/galaco/go-me-engine/valve/material"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"math"
)

type bspstructs struct {
	faces      []face.Face
	planes     []plane.Plane
	vertexes   []mgl32.Vec3
	surfEdges  []int32
	edges      [][2]uint16
	texInfos   []texinfo.TexInfo
	dispInfos  []dispinfo.DispInfo
	dispVerts  []dispvert.DispVert
	pakFile    *lumps.Pakfile
	visibility *visibility.Vis
}

func LoadMap(file *bsp.Bsp) *components.BspComponent {
	FileManager := filesystem.GetFileManager()
	bspStructure := bspstructs{
		faces:      file.GetLump(bsp.LUMP_FACES).(*lumps.Face).GetData(),
		planes:     file.GetLump(bsp.LUMP_PLANES).(*lumps.Planes).GetData(),
		vertexes:   file.GetLump(bsp.LUMP_VERTEXES).(*lumps.Vertex).GetData(),
		surfEdges:  file.GetLump(bsp.LUMP_SURFEDGES).(*lumps.Surfedge).GetData(),
		edges:      file.GetLump(bsp.LUMP_EDGES).(*lumps.Edge).GetData(),
		texInfos:   file.GetLump(bsp.LUMP_TEXINFO).(*lumps.TexInfo).GetData(),
		dispInfos:  file.GetLump(bsp.LUMP_DISPINFO).(*lumps.DispInfo).GetData(),
		dispVerts:  file.GetLump(bsp.LUMP_DISP_VERTS).(*lumps.DispVert).GetData(),
		pakFile:    file.GetLump(bsp.LUMP_PAKFILE).(*lumps.Pakfile),
		visibility: file.GetLump(bsp.LUMP_VISIBILITY).(*lumps.Visibility).GetData(),
	}

	meshList := make([]interfaces.IPrimitive, len(bspStructure.faces))
	materialList := make([]*texinfo.TexInfo, len(bspStructure.faces))

	// BSP FACES
	for idx, f := range bspStructure.faces {
		materialList[idx] = &bspStructure.texInfos[f.TexInfo]

		if f.DispInfo > -1 {
			meshList[idx] = generateDisplacementFace(&f, &bspStructure)
			// This face is a displacement instead!
		} else {
			meshList[idx] = generateBspFace(&f, &bspStructure)
		}
	}

	// Build bsp kd tree
	visibilityTree := tree.BuildTree(file, meshList)

	//MATERIALS
	// Open VPK filesystem
	vpkHandle, err := vpk.OpenVPK("data/cstrike/cstrike_pak")
	if err != nil {
		log.Fatal(err)
	}
	stringTable := stringtable.GetTable(file)
	material2.LoadMaterialList(bspStructure.pakFile, vpkHandle, stringtable.SortUnique(stringTable, materialList))

	// Add MATERIALS TO FACES
	for idx, primitive := range meshList {
		faceVmt, _ := stringTable.GetString(int(bspStructure.texInfos[bspStructure.faces[idx].TexInfo].TexData))
		vmtPath := faceVmt
		baseTexturePath := "-1"
		if FileManager.GetFile(vmtPath) != nil {
			baseTexturePath = FileManager.GetFile(vmtPath).(*material2.Vmt).GetProperty("basetexture").AsString() + ".vtf"
		}
		if FileManager.GetFile(baseTexturePath) != nil {
			mat := FileManager.GetFile(baseTexturePath).(*material.Material)
			primitive.(*base.Primitive).AddMaterial(mat)
			primitive.(*base.Primitive).AddTextureCoordinateData(texCoordsForFaceFromTexInfo(primitive.GetVertices(), &bspStructure.texInfos[bspStructure.faces[idx].TexInfo], mat.GetWidth(), mat.GetHeight()))
		} else {
			primitive.(*base.Primitive).AddTextureCoordinateData(texCoordsForFaceFromTexInfo(primitive.GetVertices(), &bspStructure.texInfos[bspStructure.faces[idx].TexInfo], 1, 1))
		}

		primitive.GenerateGPUBuffer()
	}

	return components.NewBspComponent(visibilityTree, meshList, bspStructure.visibility)
}

// Create primitives from face data in the bsp
func generateBspFace(f *face.Face, bspStructure *bspstructs) interfaces.IPrimitive {
	expF := make([]uint16, 0)
	expV := make([]float32, 0)
	expN := make([]float32, 0)

	planeNormal := bspStructure.planes[f.Planenum].Normal
	// All surfedges associated with this face
	// surfEdges are basically indices into the edges lump
	faceSurfEdges := bspStructure.surfEdges[f.FirstEdge:(f.FirstEdge + int32(f.NumEdges))]
	rootIndex := uint16(0)
	for idx, surfEdge := range faceSurfEdges {
		edge := bspStructure.edges[int(math.Abs(float64(surfEdge)))]
		e1 := 0
		e2 := 1
		if surfEdge < 0 {
			e1 = 1
			e2 = 0
		}
		//Capture root indice
		if idx == 0 {
			rootIndex = edge[e1]
		} else {
			// Just create a triangle for every edge now
			expF = append(expF, rootIndex, edge[e1], edge[e2])
			expV = append(expV, bspStructure.vertexes[rootIndex].X(), bspStructure.vertexes[rootIndex].Y(), bspStructure.vertexes[rootIndex].Z())
			expN = append(expN, planeNormal.X(), planeNormal.Y(), planeNormal.Z())
			expV = append(expV, bspStructure.vertexes[edge[e1]].X(), bspStructure.vertexes[edge[e1]].Y(), bspStructure.vertexes[edge[e1]].Z())
			expN = append(expN, planeNormal.X(), planeNormal.Y(), planeNormal.Z())
			expV = append(expV, bspStructure.vertexes[edge[e2]].X(), bspStructure.vertexes[edge[e2]].Y(), bspStructure.vertexes[edge[e2]].Z())
			expN = append(expN, planeNormal.X(), planeNormal.Y(), planeNormal.Z())
		}
	}

	return base.NewPrimitive(expV, expF, expN)
}

// Create Primitives from Displacement faces tied to faces
// in the bsp
// @TODO implement me
func generateDisplacementFace(f *face.Face, bspStructure *bspstructs) interfaces.IPrimitive {
	//numSubDivisions := int(disp.Power*disp.Power)
	//numVerts := numSubDivisions * numSubDivisions
	//dispVertList := (*dispVerts)[disp.DispVertStart:disp.DispVertStart + int32(numVerts)]
	//
	//// Construct a subdivided vertex positions
	//for x := 0; x < numSubDivisions; x++ {
	//	for y := 0; y < numSubDivisions; y++ {
	//		log.Println(dispVertList[x + y].Dist)
	//	}
	//}
	return generateBspFace(f, bspStructure)
}

// Generate texturecoordinates for face data
func texCoordsForFaceFromTexInfo(vertexes []float32, tx *texinfo.TexInfo, width int, height int) (uvs []float32) {
	for idx := 0; idx < len(vertexes); idx += 3 {
		//u = tv0,0 * x + tv0,1 * y + tv0,2 * z + tv0,3
		u := ((tx.TextureVecsTexelsPerWorldUnits[0][0] * vertexes[idx]) +
			(tx.TextureVecsTexelsPerWorldUnits[0][1] * vertexes[idx+1]) +
			(tx.TextureVecsTexelsPerWorldUnits[0][2] * vertexes[idx+2]) +
			tx.TextureVecsTexelsPerWorldUnits[0][3]) / float32(width)

		//v = tv1,0 * x + tv1,1 * y + tv1,2 * z + tv1,3
		v := ((tx.TextureVecsTexelsPerWorldUnits[1][0] * vertexes[idx]) +
			(tx.TextureVecsTexelsPerWorldUnits[1][1] * vertexes[idx+1]) +
			(tx.TextureVecsTexelsPerWorldUnits[1][2] * vertexes[idx+2]) +
			tx.TextureVecsTexelsPerWorldUnits[1][3]) / float32(height)

		uvs = append(uvs, u, v)
	}

	return uvs
}
