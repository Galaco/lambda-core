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
	vbo uint32
	vao uint32
	indicesBuffer uint32
	uvBuffer uint32
	faceMode uint32
	material material.IMaterial
	isBoundToGPU bool
}

func (primitive *Primitive) Bind() {
	gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(primitive.vao)

	// UV's
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, primitive.uvBuffer)
	gl.VertexAttribPointer(
		1,                // The attribute we want to configure
		2,                 // size : U+V => 2
		gl.FLOAT,               // type
		false,		// normalized?
		0,               // stride
		nil)        		// array buffer offset
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
	if primitive.isBoundToGPU == true {
		return
	}
	// Gen vbo data
	gl.GenBuffers(1, &primitive.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, primitive.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(primitive.vertices), gl.Ptr(primitive.vertices), gl.STATIC_DRAW)

	// gen vao data
	gl.GenVertexArrays(1, &primitive.vao)
	gl.BindVertexArray(primitive.vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, primitive.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// gen uv data
	gl.GenBuffers(1, &primitive.uvBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, primitive.uvBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(primitive.textureCoordinates) * 2, gl.Ptr(primitive.textureCoordinates), gl.STATIC_DRAW)


	//gl.GenBuffers(1, &primitive.indicesBuffer)
	//gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, primitive.indicesBuffer)
	//gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(primitive.indices) * 4, gl.Ptr(primitive.indices), gl.STATIC_DRAW)
	//
	//gl.EnableVertexAttribArray(1)
	//gl.BindBuffer(gl.ARRAY_BUFFER, primitive.uvBuffer)
	//gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	//gl.BufferData(gl.ARRAY_BUFFER, len(primitive.textureCoordinates) * 4, gl.Ptr(primitive.textureCoordinates), gl.STATIC_DRAW)

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
		isBoundToGPU: false,
	}
}