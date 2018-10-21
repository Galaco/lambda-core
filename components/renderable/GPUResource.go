package renderable

import (
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
)

// A collection of renderable primitives/submeshes
type GPUResource struct {
	primitives   []primitive.IPrimitive
	isBoundToGPU bool
}

// Ensure that data contained is passed to the GPU
func (resource *GPUResource) Prepare() {
	if resource.isBoundToGPU == true {
		return
	}
	for _, p := range resource.primitives {
		p.GenerateGPUBuffer()
	}
}

// Add a new primitive
func (resource *GPUResource) AddPrimitive(primitive primitive.IPrimitive) {
	if resource.isBoundToGPU == true {
		primitive.GenerateGPUBuffer()
	}
	resource.primitives = append(resource.primitives, primitive)
}

func (resource *GPUResource) AddPrimitives(primitives []primitive.IPrimitive) {
	for _, p := range primitives {
		resource.AddPrimitive(p)
	}
}

// Get all primitives/submeshes
func (resource *GPUResource) GetPrimitives() []primitive.IPrimitive {
	return resource.primitives
}

func NewGPUResource(primitives []primitive.IPrimitive) *GPUResource {
	return &GPUResource{
		primitives:   primitives,
		isBoundToGPU: false,
	}
}

type GPUResourceDynamic struct {
	GPUResource
}

func (resource *GPUResourceDynamic) Reset() {
	resource.primitives = []primitive.IPrimitive{}
	resource.isBoundToGPU = false
}

func NewGPUResourceDynamic(primitives []primitive.IPrimitive) *GPUResourceDynamic {
	return &GPUResourceDynamic{
		GPUResource: GPUResource{
			primitives:   primitives,
			isBoundToGPU: false,
		},
	}
}
