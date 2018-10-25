package primitive

import (
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/go-gl/gl/v4.1-core/gl"
)

var cubeVerts = []float32{
	-1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, -1.0, 1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	1.0, 1.0, -1.0,
	1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0,
	1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,
	1.0, -1.0, 1.0,
	1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0,
	-1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,
	1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,
	1.0, 1.0, -1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, -1.0,
	-1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0,
	1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, 1.0, -1.0,
	1.0, -1.0, 1.0,
	1.0, 1.0, -1.0,
	1.0, 1.0, 1.0,
}

var cubeNormals = cubeVerts

var cubeUVs = []float32{
	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,

	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,

	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,

	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 0,
	0, 1,
	1, 1,
}

type Cube struct {
	mesh.Mesh
}

func (cube *Cube) GetFaceMode() uint32 {
	return gl.TRIANGLES
}

func NewCube() *Cube {
	c := &Cube{
		Mesh: *mesh.NewMesh(),
	}
	c.AddVertex(cubeVerts...)
	c.AddNormal(cubeNormals...)
	c.AddTextureCoordinate(cubeUVs...)
	c.Finish()

	return c
}