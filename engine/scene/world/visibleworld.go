package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/model"
)

type VisibleWorld struct {
	entity.Base

	world        *model.Bsp
	visibleProps []*StaticProp
	sky          *Sky
}

func (entity *VisibleWorld) Bsp() *model.Bsp {
	return entity.world
}

func (entity *VisibleWorld) Staticprops() []*StaticProp {
	return entity.visibleProps
}

func (entity *VisibleWorld) Sky() *Sky {
	return entity.sky
}

func NewVisibleWorld() *VisibleWorld {
	c := VisibleWorld{
		world:        model.NewBsp(mesh.NewMesh()),
		visibleProps: make([]*StaticProp, 0),
	}

	return &c
}
