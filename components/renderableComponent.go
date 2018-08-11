package components

import (
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/components/renderable"
)



type RenderableComponent struct {
	base.Component
	renderable *renderable.GPUResource
}

func (component *RenderableComponent) Initialize() {
}

func (component *RenderableComponent) SetRenderableResource(resource *renderable.GPUResource) {
	component.renderable = resource
}

func (component *RenderableComponent) GetRenderable() *renderable.GPUResource {
	return component.renderable
}


func NewRenderableComponent() *RenderableComponent{
	c := &RenderableComponent{}
	c.Etype = T_RenderableComponent

	return c
}