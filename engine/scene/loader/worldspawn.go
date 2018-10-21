package loader

import (
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/material"
	material2 "github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
	sceneVisibility "github.com/galaco/Gource-Engine/engine/scene/visibility"
	"github.com/galaco/Gource-Engine/entity"
	"github.com/galaco/Gource-Engine/lib/stringtable"
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/dispinfo"
	"github.com/galaco/bsp/primitives/dispvert"
	"github.com/galaco/bsp/primitives/face"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/bsp/primitives/visibility"
	"github.com/go-gl/mathgl/mgl32"
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
	game       *lumps.Game
}

func LoadMap(file *bsp.Bsp) *entity.WorldSpawn {
	ResourceManager := filesystem.Manager()
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
		game:       file.GetLump(bsp.LUMP_GAME_LUMP).(*lumps.Game),
	}

	meshList := make([]primitive.IPrimitive, len(bspStructure.faces))
	materialList := make([]*texinfo.TexInfo, len(bspStructure.faces))

	// BSP FACES
	for idx, f := range bspStructure.faces {
		materialList[idx] = &bspStructure.texInfos[f.TexInfo]

		if f.DispInfo > -1 {
			// This face is a displacement
			meshList[idx] = generateDisplacementFace(&f, &bspStructure)
		} else {
			meshList[idx] = generateBspFace(&f, &bspStructure)
		}
	}

	//MATERIALS
	stringTable := stringtable.GetTable(file)
	material2.LoadMaterialList(stringtable.SortUnique(stringTable, materialList))

	// Add MATERIALS TO FACES
	for idx, mesh := range meshList {
		if mesh == nil {
			continue
		}
		faceVmt, _ := stringTable.GetString(int(bspStructure.texInfos[bspStructure.faces[idx].TexInfo].TexData))
		vmtPath := faceVmt
		baseTexturePath := "-1"
		if ResourceManager.Has(vmtPath) {
			baseTexturePath = ResourceManager.Get(vmtPath).(*material2.Vmt).GetProperty("basetexture").AsString() + ".vtf"
		}
		if ResourceManager.Has(baseTexturePath) {
			mat := ResourceManager.Get(baseTexturePath).(*material.Material)
			mesh.(*primitive.Primitive).AddMaterial(mat)
			mesh.(*primitive.Primitive).AddTextureCoordinateData(texCoordsForFaceFromTexInfo(mesh.GetVertices(), &bspStructure.texInfos[bspStructure.faces[idx].TexInfo], mat.GetWidth(), mat.GetHeight()))
		} else {
			mesh.(*primitive.Primitive).AddTextureCoordinateData(texCoordsForFaceFromTexInfo(mesh.GetVertices(), &bspStructure.texInfos[bspStructure.faces[idx].TexInfo], 1, 1))
		}

		mesh.GenerateGPUBuffer()
	}

	// Load static props
	LoadStaticProps(bspStructure.game.GetStaticPropLump())

	visData := sceneVisibility.NewVisFromBSP(file)

	return entity.NewWorld(meshList, visData)
}

// Create primitives from face data in the bsp
func generateBspFace(f *face.Face, bspStructure *bspstructs) primitive.IPrimitive {
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

	return primitive.NewPrimitive(expV, expF, expN)
}

// Create Primitive from Displacement face
// This is based on:
// https://github.com/Metapyziks/VBspViewer/blob/master/Assets/VBspViewer/Scripts/Importing/VBsp/VBspFile.cs
func generateDisplacementFace(f *face.Face, bspStructure *bspstructs) primitive.IPrimitive {
	corners := make([]mgl32.Vec3, 4)
	normal := bspStructure.planes[f.Planenum].Normal

	info := bspStructure.dispInfos[f.DispInfo]
	size := int(1 << uint32(info.Power))
	firstCorner := int32(0)
	firstCornerDist2 := float32(math.MaxFloat32)

	for surfId := f.FirstEdge; surfId < f.FirstEdge+int32(f.NumEdges); surfId++ {
		surfEdge := bspStructure.surfEdges[surfId]
		edgeIndex := int32(math.Abs(float64(surfEdge)))
		edge := bspStructure.edges[edgeIndex]
		vert := bspStructure.vertexes[edge[0]]
		if surfEdge < 0 {
			vert = bspStructure.vertexes[edge[1]]
		}

		corners[surfId-f.FirstEdge] = vert

		dist2tmp := info.StartPosition.Sub(vert)
		dist2 := (dist2tmp.X() * dist2tmp.Y()) + (dist2tmp.Y() * dist2tmp.Y()) + (dist2tmp.Z() * dist2tmp.Z())
		if dist2 < firstCornerDist2 {
			firstCorner = surfId - f.FirstEdge
			firstCornerDist2 = dist2
		}
	}

	verts := make([]float32, 0)
	normals := make([]float32, 0)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			a := generateDispVert(int(info.DispVertStart), x, y, size, corners, firstCorner, &bspStructure.dispVerts)
			b := generateDispVert(int(info.DispVertStart), x, y+1, size, corners, firstCorner, &bspStructure.dispVerts)
			c := generateDispVert(int(info.DispVertStart), x+1, y+1, size, corners, firstCorner, &bspStructure.dispVerts)
			d := generateDispVert(int(info.DispVertStart), x+1, y, size, corners, firstCorner, &bspStructure.dispVerts)

			// Split into triangles
			verts = append(verts, a.X(), a.Y(), a.Z(), b.X(), b.Y(), b.Z(), c.X(), c.Y(), c.Z())
			normals = append(normals, normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z())
			verts = append(verts, a.X(), a.Y(), a.Z(), c.X(), c.Y(), c.Z(), d.X(), d.Y(), d.Z())
			normals = append(normals, normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z())
		}
	}

	return primitive.NewPrimitive(verts, make([]uint16, 3), normals)
}

func generateDispVert(offset int, x int, y int, size int, corners []mgl32.Vec3, firstCorner int32, dispVerts *[]dispvert.DispVert) mgl32.Vec3 {
	vert := (*dispVerts)[offset+x+y*(size+1)]

	tx := float32(x / size)
	ty := float32(y / size)
	sx := 1 - tx
	sy := 1 - ty

	cornerA := corners[(0+firstCorner)&3]
	cornerB := corners[(1+firstCorner)&3]
	cornerC := corners[(2+firstCorner)&3]
	cornerD := corners[(3+firstCorner)&3]

	origin := ((cornerB.Mul(sx).Add(cornerC.Mul(tx))).Mul(ty)).Add((cornerA.Mul(sx).Add(cornerD.Mul(tx))).Mul(sy))

	return origin.Add(vert.Vec.Mul(vert.Dist))
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
