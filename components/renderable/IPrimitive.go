package renderable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/galaco/go-me-engine/components/renderable/material"
)

type IPrimitive interface {
	GetVertices() []float32
	GetIndices() []uint16
	GetTextureCoordinates() []float32
	GetMaterial() material.IMaterial
	GetFaceMode() uint32
	GenerateGPUBuffer()
	Bind()
}

type Primitive struct {
	vertices []float32
	indices []uint16
	textureCoordinates []float32
	buffer uint32
	faceMode uint32
	material material.IMaterial
}

func (primitive *Primitive) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, primitive.buffer)
}

func (primitive *Primitive) GetFaceMode() uint32 {
	return primitive.faceMode
}

func (primitive *Primitive) GetVertices() []float32 {
	return primitive.vertices
}

func (primitive *Primitive) GetIndices() []uint16 {
	return primitive.indices
}
func (primitive *Primitive) GetTextureCoordinates() []float32 {
	return primitive.textureCoordinates
}

func (primitive *Primitive) GetMaterial() material.IMaterial {
	return primitive.material
}

func (primitive *Primitive) AddVertexData(vertices []float32) {
	primitive.vertices = vertices
}

func (primitive *Primitive) AddIndexData(indices []uint16) {
	primitive.indices = indices
}

func (primitive *Primitive) AddTextureCoordinateData(textureCoordinates []float32) {
	primitive.textureCoordinates = textureCoordinates
}

func (primitive *Primitive) AddMaterial(material material.IMaterial) {
	primitive.material = material
}

func (primitive *Primitive) GenerateGPUBuffer() {
	gl.GenBuffers(1, &primitive.buffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, primitive.buffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(primitive.indices) * 4, gl.Ptr(primitive.indices), gl.STATIC_DRAW)

	switch len(primitive.indices) {
	case 1:
		primitive.faceMode = gl.POINTS
	case 2:
		primitive.faceMode = gl.LINES
	case 3:
		primitive.faceMode = gl.TRIANGLES
	case 4:
		primitive.faceMode = gl.QUADS
	default:
		primitive.faceMode = gl.TRIANGLES
	}
}

func NewPrimitive(vertices []float32, indices []uint16) *Primitive {
	return &Primitive{
		vertices: vertices,
		indices: indices,
	}
}