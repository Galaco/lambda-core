package components

import (
	"github.com/galaco/go-me-engine/components/renderable/material"
)

type Skybox struct {
	RenderableComponent
	cubemap *material.Cubemap
}

func (skybox *Skybox) GetCubemap() *material.Cubemap{
	return skybox.cubemap
}

func NewSkybox(cubemap *material.Cubemap) *Skybox {
	return &Skybox{
		cubemap: cubemap,
	}
}