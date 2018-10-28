package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
)

type VisibleWorld struct {
	entity.Base
	visibleModel *model.Model
	visibleProps []*StaticProp
	sky          *Sky
}

func (entity *VisibleWorld) Bsp() *model.Model {
	return entity.visibleModel
}

func (entity *VisibleWorld) Staticprops() []*StaticProp {
	return entity.visibleProps
}

func (entity *VisibleWorld) Sky() *Sky {
	return entity.sky
}

func NewVisibleWorld() *VisibleWorld {
	c := VisibleWorld{
		visibleModel: model.NewModel("worldspawn_visible"),
		visibleProps: make([]*StaticProp, 0),
	}

	return &c
}
