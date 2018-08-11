package renderable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)



type GPUResource struct {
	isBound bool
	vao uint32
	vertexData []float32
}

func (resource *GPUResource) BindData() {
	if resource.isBound == false {
		resource.vao = resource.generateVao(resource.vertexData)
		resource.isBound = true
	}
}

func (resource *GPUResource) GetVertexData() []float32 {
	return resource.vertexData
}

func (resource *GPUResource) GetVao() uint32 {
	return resource.vao
}

func (resource *GPUResource) generateVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func NewGPUResource(vertexData []float32) *GPUResource {
	return &GPUResource{
		vertexData: vertexData,
	}
}