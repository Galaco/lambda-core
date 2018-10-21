package components

import (
	"github.com/galaco/Gource-Engine/engine/component"
	"github.com/galaco/Gource-Engine/engine/mesh"
)

type RenderableComponent struct {
	component.Component
	renderables []mesh.IMesh
}

func (component *RenderableComponent) AddRenderableResource(resource mesh.IMesh) {
	// Ensure our GPU resource is ready to use
	resource.Prepare()
	component.renderables = append(component.renderables, resource)
}

// Return a list of all renderable from this component
// this can be many different collections of primitives
func (component *RenderableComponent) GetRenderables() []mesh.IMesh {
	return component.renderables
}

func NewRenderableComponent() *RenderableComponent {
	c := RenderableComponent{}

	return &c
}
