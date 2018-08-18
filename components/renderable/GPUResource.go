package renderable

// A collection of renderable primitives/submeshes
type GPUResource struct {
	primitives []IPrimitive
	isBoundToGPU bool
}

// Ensure that data contained is passed to the GPU
func (resource *GPUResource) Prepare() {
	if resource.isBoundToGPU == true {
		return
	}
	for _,p := range resource.primitives {
		p.GenerateGPUBuffer()
	}
}

// Add a new primitive
func (resource *GPUResource) AddPrimitive(primitive IPrimitive) {
	resource.primitives = append(resource.primitives, primitive)
	resource.isBoundToGPU = false
}

// Get all primitives/submeshes
func (resource *GPUResource) GetPrimitives() []IPrimitive {
	return resource.primitives
}

func NewGPUResource(primitives []IPrimitive) *GPUResource {
	return &GPUResource{
		primitives: primitives,
		isBoundToGPU: false,
	}
}