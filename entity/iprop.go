package entity

import "github.com/galaco/Gource-Engine/engine/model"

type IProp interface {
	GetModel() *model.Model
	SetModel(model *model.Model)
}

type PropBase struct {
	model *model.Model
}

func (prop *PropBase) SetModel(model *model.Model) {
	prop.model = model
}

func (prop *PropBase) GetModel() *model.Model {
	return prop.model
}
