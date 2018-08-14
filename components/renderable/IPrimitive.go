package renderable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type IPrimitive interface {
	GetVertices() []float32
	GetIndices() []uint16
	GenerateGPUBuffer()
	Bind()
}

type Primitive struct {
	vertices []float32
	indices []uint16
	buffer uint32
}

func (primitive *Primitive) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, primitive.buffer)
}

func (primitive *Primitive) GetVertices() []float32 {
	return primitive.vertices
}

func (primitive *Primitive) GetIndices() []uint16 {
	return primitive.indices
}

func (primitive *Primitive) AddVertexData(vertices []float32) {
	primitive.vertices = vertices
}

func (primitive *Primitive) AddIndexData(indices []uint16) {
	primitive.indices = indices
}

func (primitive *Primitive) GenerateGPUBuffer() {
	gl.GenBuffers(1, &primitive.buffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, primitive.buffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(primitive.indices) * 4, gl.Ptr(primitive.indices), gl.STATIC_DRAW)
}

func NewPrimitive(vertices []float32, indices []uint16) *Primitive {
	return &Primitive{
		vertices: vertices,
		indices: indices,
	}
}