package common

import (
	"github.com/galaco/Lambda-Core/core/entity"
	entity2 "github.com/galaco/Lambda-Core/game/entity"
)

type PropDynamicOrnament struct {
	entity.Base
	entity2.PropBase
}

func (entity *PropDynamicOrnament) New() entity.IEntity {
	return &PropDynamicOrnament{}
}

func (entity PropDynamicOrnament) Classname() string {
	return "prop_dynamic_ornament"
}
