package model

import (
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
)

// A collection of renderable primitives/submeshes
type Model struct {
	primitives   []primitive.IPrimitive
	isBoundToGPU bool
}

// Ensure that data contained is passed to the GPU
func (resource *Model) Prepare() {
	if resource.isBoundToGPU == true {
		return
	}
	for _, p := range resource.primitives {
		p.GenerateGPUBuffer()
	}
}

// Add a new primitive
func (resource *Model) AddPrimitive(primitive primitive.IPrimitive) {
	if resource.isBoundToGPU == true {
		primitive.GenerateGPUBuffer()
	}
	resource.primitives = append(resource.primitives, primitive)
}

func (resource *Model) AddPrimitives(primitives []primitive.IPrimitive) {
	for _, p := range primitives {
		resource.AddPrimitive(p)
	}
}

// Get all primitives/submeshes
func (resource *Model) GetPrimitives() []primitive.IPrimitive {
	return resource.primitives
}

func (resource *Model) Reset() {
	resource.primitives = []primitive.IPrimitive{}
	resource.isBoundToGPU = false
}

func NewModel(primitives []primitive.IPrimitive) *Model {
	return &Model{
		primitives:   primitives,
		isBoundToGPU: false,
	}
}