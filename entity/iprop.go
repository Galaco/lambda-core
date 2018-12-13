package entity

import "github.com/galaco/Gource-Engine/engine/model"

// IProp Base renderable prop interface
type IProp interface {
	GetModel() *model.Model
	SetModel(model *model.Model)
}

// PropBase is a minimal renderable prop entity
type PropBase struct {
	model *model.Model
}

func (prop *PropBase) SetModel(model *model.Model) {
	prop.model = model
}

func (prop *PropBase) GetModel() *model.Model {
	return prop.model
}
