package renderable

import "github.com/galaco/go-me-engine/engine/interfaces"

// A collection of renderable primitives/submeshes
type GPUResource struct {
	primitives   []interfaces.IPrimitive
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
func (resource *GPUResource) AddPrimitive(primitive interfaces.IPrimitive) {
	resource.primitives = append(resource.primitives, primitive)
	resource.isBoundToGPU = false
}

func (resource *GPUResource) AddPrimitives(primitives []interfaces.IPrimitive) {
	resource.primitives = append(resource.primitives, primitives...)
	resource.isBoundToGPU = false
}

// Get all primitives/submeshes
func (resource *GPUResource) GetPrimitives() []interfaces.IPrimitive {
	return resource.primitives
}

func NewGPUResource(primitives []interfaces.IPrimitive) *GPUResource {
	return &GPUResource{
		primitives:   primitives,
		isBoundToGPU: false,
	}
}

type GPUResourceDynamic struct {
	GPUResource
}

func (resource *GPUResourceDynamic) Reset() {
	resource.primitives = []interfaces.IPrimitive{}
	resource.isBoundToGPU = false
}

func NewGPUResourceDynamic(primitives []interfaces.IPrimitive) *GPUResourceDynamic {
	return &GPUResourceDynamic{
		GPUResource: GPUResource{
			primitives:   primitives,
			isBoundToGPU: false,
		},
	}
}
