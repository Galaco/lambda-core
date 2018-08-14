package bsp

import (
	"math"
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/primitives/face"
	"github.com/galaco/bsp/lumps"
	"github.com/go-gl/mathgl/mgl32"
)

func GenerateTrianglesFromBSP(file *bsp.Bsp) ([]float32, []uint16) {
	var expVerts []float32
	var expIndices []uint16

	fl := *file.GetLump(bsp.LUMP_FACES).GetContents()
	faces := (fl).(lumps.Face).GetData().(*[]face.Face)

	vs := *file.GetLump(bsp.LUMP_VERTEXES).GetContents()
	vertexes := (vs).(lumps.Vertex).GetData().(*[]mgl32.Vec3)

	sf := *file.GetLump(bsp.LUMP_SURFEDGES).GetContents()
	surfEdges := (sf).(lumps.Surfedge).GetData().(*[]int32)

	ed := *file.GetLump(bsp.LUMP_EDGES).GetContents()
	edges := (ed).(lumps.Edge).GetData().(*[][2]uint16)


	for _,v := range *vertexes {
		expVerts = append(expVerts, v.X(), v.Y(), v.Z())
	}

	for _,f := range *faces {
		//// Just so we render triangles

		// rewrite to just get a flat vertex array, then create indices into the vertex data for the quads
		surfEdges := (*surfEdges)[f.FirstEdge:(f.FirstEdge+int32(f.NumEdges))]
		for _,surfEdge := range surfEdges {
			edge := (*edges)[int(math.Abs(float64(surfEdge)))]

			// Fix reverse ordering
			if surfEdge >= 0 {
				expIndices = append(expIndices, edge[0], edge[1])
			} else {
				expIndices = append(expIndices, edge[0], edge[1])
			}
		}
	}

	return expVerts,expIndices
}

