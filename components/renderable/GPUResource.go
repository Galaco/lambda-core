package renderable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type GPUResource struct {
	vbo        uint32
	vao        uint32
	vertices   []float32
	primitives []IPrimitive
}

func (resource *GPUResource) Bind() {
	gl.BindVertexArray(resource.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(resource.vertices) / 3))
	//gl.BindBuffer(gl.ARRAY_BUFFER, resource.vbo)
	//gl.EnableVertexAttribArray(0)
	//gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func (resource *GPUResource) AddPrimitive(primitive IPrimitive) {
	primitive.GenerateGPUBuffer()
	resource.primitives = append(resource.primitives, primitive)
}

func (resource *GPUResource) GetPrimitives() []IPrimitive {
	return resource.primitives
}

func (resource *GPUResource) GetVertexData() []float32 {
	return resource.vertices
}

func (resource *GPUResource) GenerateGPUBuffer() {
	// Gen vbo data
	gl.GenBuffers(1, &resource.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, resource.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(resource.vertices), gl.Ptr(resource.vertices), gl.STATIC_DRAW)

	// gen vao data
	gl.GenVertexArrays(1, &resource.vao)
	gl.BindVertexArray(resource.vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, resource.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func NewGPUResource(vertices []float32) *GPUResource {
	return &GPUResource{
		vertices: vertices,
	}
}