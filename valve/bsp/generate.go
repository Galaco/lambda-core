package bsp

import (
	"math"
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/primitives/face"
	"github.com/galaco/bsp/lumps"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/bsp/primitives/dispinfo"
	"github.com/galaco/bsp/primitives/dispvert"
	"log"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	vpk2 "github.com/galaco/vpk2"
	"github.com/galaco/go-me-engine/valve/libwrapper/vpk"
	"github.com/galaco/go-me-engine/valve/libwrapper/stringtable"
	"github.com/galaco/go-me-engine/valve/libwrapper/vtf"
)

type bspstructs struct {
	faces []face.Face
	planes []plane.Plane
	vertexes []mgl32.Vec3
	surfEdges []int32
	edges [][2]uint16
	texInfos []texinfo.TexInfo
}

func LoadMap(file *bsp.Bsp) ([]interfaces.IPrimitive) {
	FileManager := filesystem.GetFileManager()
	bspStructure := bspstructs{
		faces:     *(*file.GetLump(bsp.LUMP_FACES).GetContents()).(lumps.Face).GetData().(*[]face.Face),
		planes:    (*file.GetLump(bsp.LUMP_PLANES).GetContents()).(lumps.Planes).GetData().([]plane.Plane),
		vertexes:  *(*file.GetLump(bsp.LUMP_VERTEXES).GetContents()).(lumps.Vertex).GetData().(*[]mgl32.Vec3),
		surfEdges: *(*file.GetLump(bsp.LUMP_SURFEDGES).GetContents()).(lumps.Surfedge).GetData().(*[]int32),
		edges:     *(*file.GetLump(bsp.LUMP_EDGES).GetContents()).(lumps.Edge).GetData().(*[][2]uint16),
		texInfos:  *(*file.GetLump(bsp.LUMP_TEXINFO).GetContents()).(lumps.TexInfo).GetData().(*[]texinfo.TexInfo),
	}

	meshList := make([]interfaces.IPrimitive, len(bspStructure.faces))
	materialList := []*texinfo.TexInfo{}

	// BSP FACES
	for idx,f := range bspStructure.faces {
		materialList = append(materialList, &bspStructure.texInfos[f.TexInfo])

		if f.DispInfo > -1 {
			//meshList[idx] = genereateDisplacementFace(file, &f)
			// This face is a displacement instead!
		} else {
			meshList[idx] = generateBspFace(&f, &bspStructure)
		}
	}

	//MATERIALS
	// Open VPK filesystem
	vpkHandle,err := vpk.OpenVPK("data/cstrike/cstrike_pak")
	if err != nil {
		log.Fatal(err)
	}
	stringTable := stringtable.GetTable(file)
	loadMaterials(vpkHandle, stringtable.SortUnique(stringTable, materialList))

	// Add MATERIALS TO FACES
	for idx,primitive := range meshList {
		if primitive == nil {
			continue
		}
		target,_ := stringTable.GetString(int(bspStructure.texInfos[bspStructure.faces[idx].TexInfo].TexData))
		if FileManager.GetFile(target) != nil {
			mat := FileManager.GetFile(target).(*material.Material)
			primitive.(*base.Primitive).AddMaterial(mat)
			primitive.(*base.Primitive).AddTextureCoordinateData(texCoordsForFaceFromTexInfo(primitive.GetVertices(), &bspStructure.texInfos[bspStructure.faces[idx].TexInfo], mat.GetWidth(), mat.GetHeight()))
		} else {
			primitive.(*base.Primitive).AddTextureCoordinateData(texCoordsForFaceFromTexInfo(primitive.GetVertices(), &bspStructure.texInfos[bspStructure.faces[idx].TexInfo], 1, 1))
		}
	}


	return meshList
}

func generateBspFace(f *face.Face, bspStructure *bspstructs) interfaces.IPrimitive{
	expF := make([]uint16, 0)
	expV := make([]float32, 0)
	expN := make([]float32, 0)

	planeNormal := bspStructure.planes[f.Planenum].Normal
	// All surfedges associated with this face
	// surfEdges are basically indices into the edges lump
	faceSurfEdges := bspStructure.surfEdges[f.FirstEdge:(f.FirstEdge+int32(f.NumEdges))]
	rootIndex := uint16(0)
	for idx,surfEdge := range faceSurfEdges {
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

func genereateDisplacementFace(file *bsp.Bsp, referenceFace *face.Face) {//(expVerts [][]float32, expIndices [][]uint16, expTexInfos []texinfo.TexInfo, expNormals [][]float32) {
	dispInfos := (*file.GetLump(bsp.LUMP_DISPINFO).GetContents()).(lumps.DispInfo).GetData().(*[]dispinfo.DispInfo)
	dispVerts := (*file.GetLump(bsp.LUMP_DISP_VERTS).GetContents()).(lumps.DispVert).GetData().(*[]dispvert.DispVert)

	disp := (*dispInfos)[referenceFace.DispInfo]

	numSubDivisions := int(disp.Power*disp.Power)
	numVerts := numSubDivisions * numSubDivisions
	dispVertList := (*dispVerts)[disp.DispVertStart:disp.DispVertStart + int32(numVerts)]

	// Construct a subdivided vertex positions
	for x := 0; x < numSubDivisions; x++ {
		for y := 0; y < numSubDivisions; y++ {
			log.Println(dispVertList[x + y].Dist)
		}
	}
}

func loadMaterials(vpkHandle *vpk2.VPK, materialList []string) {
	FileManager := filesystem.GetFileManager()

	for _,materialPath := range materialList {
		// Load file from vpk into memory
		vpkFile := vpkHandle.Entry("materials/" + materialPath + ".vtf")
		if vpkFile == nil {
			log.Println("Couldnt find material: materials/" + materialPath + ".vtf")
			continue
		}
		file,err := vpkFile.Open()

		// Its quite possible for a texture to be missing, just skip it.
		if err != nil {
			continue
		}

		// Attempt to parse the vtf into color data we can use,
		// if this fails (it shouldn't) we can treat it like it was missing
		texture,err := vtf.ReadFromStream(file)
		if err != nil {
			log.Println(err)
			continue
		}
		// Store file containing raw data in memory
		FileManager.AddFile(
			material.NewMaterial(
				materialPath,
				texture,
				int(texture.GetHeader().Width),
				int(texture.GetHeader().Height)))
		// Finally generate the gpu buffer for the material
		FileManager.GetFile(materialPath).(*material.Material).GenerateGPUBuffer()
	}
}