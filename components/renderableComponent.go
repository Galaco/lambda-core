package components

import (
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/components/renderable"
)



type RenderableComponent struct {
	base.Component
	renderables []*renderable.GPUResource
}

func (component *RenderableComponent) Initialize() {
}

func (component *RenderableComponent) AddRenderableResource(resource *renderable.GPUResource) {
	// Ensure our GPU resource is ready to use
	resource.Prepare()
	component.renderables = append(component.renderables, resource)
}

func (component *RenderableComponent) GetRenderables() []*renderable.GPUResource {
	return component.renderables
}


func NewRenderableComponent() *RenderableComponent{
	c := RenderableComponent{}
	c.Etype = T_RenderableComponent

	return &c
}