package components

import (
	"github.com/galaco/Gource-Engine/engine/base"
	"github.com/galaco/Gource-Engine/engine/interfaces"
)

type RenderableComponent struct {
	base.Component
	renderables []interfaces.IGPUMesh
}

func (component *RenderableComponent) AddRenderableResource(resource interfaces.IGPUMesh) {
	// Ensure our GPU resource is ready to use
	resource.Prepare()
	component.renderables = append(component.renderables, resource)
}

// Return a list of all renderable from this component
// this can be many different collections of primitives
func (component *RenderableComponent) GetRenderables() []interfaces.IGPUMesh {
	return component.renderables
}

func NewRenderableComponent() *RenderableComponent {
	c := RenderableComponent{}

	return &c
}
