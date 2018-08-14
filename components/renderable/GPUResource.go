package renderable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type GPUResource struct {
	buffer uint32
	vertices []float32
	primitives []IPrimitive
}

func (resource *GPUResource) Bind() {
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, resource.buffer)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
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
	gl.GenBuffers(1, &resource.buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, resource.buffer)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(resource.vertices), gl.Ptr(resource.vertices), gl.STATIC_DRAW)
}

func NewGPUResource(vertices []float32) *GPUResource {
	return &GPUResource{
		vertices: vertices,
	}
}