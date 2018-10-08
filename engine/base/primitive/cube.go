package primitive

import (
	"github.com/galaco/Gource/engine/base"
	"github.com/galaco/Gource/engine/base/material"
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

var cubeIndices = []uint16{}

var cubeNormals = cubeVerts

var cubeUVs = []float32{
	0, 0,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 1,
	0, 0,
	1, 0,
	0, 1,
	1, 1,
}

type Cube struct {
	base.Primitive
}

func (cube *Cube) GetFaceMode() uint32 {
	return gl.TRIANGLES
}

func NewCube() *Cube {
	c := &Cube{
		*base.NewPrimitive(cubeVerts, cubeIndices, cubeNormals),
	}

	c.AddTextureCoordinateData(cubeUVs)
	mat := material.NewMaterial("placeholder", 1, 1, []uint8{255, 255, 255})
	mat.GenerateGPUBuffer()
	c.AddMaterial(mat)
	c.GenerateGPUBuffer()

	return c
}
