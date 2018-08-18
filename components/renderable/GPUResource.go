package renderable

type GPUResource struct {
	primitives []IPrimitive
}

func (resource *GPUResource) Prepare() {
	for _,p := range resource.primitives {
		p.GenerateGPUBuffer()
	}
}

func (resource *GPUResource) AddPrimitive(primitive IPrimitive) {
	resource.primitives = append(resource.primitives, primitive)
}

func (resource *GPUResource) GetPrimitives() []IPrimitive {
	return resource.primitives
}

func NewGPUResource(primitives []IPrimitive) *GPUResource {
	return &GPUResource{
		primitives: primitives,
	}
}