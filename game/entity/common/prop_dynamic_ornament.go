package common

import (
	"github.com/galaco/Gource-Engine/core/entity"
	entity2 "github.com/galaco/Gource-Engine/game/entity"
)

type PropDynamicOrnament struct {
	entity.Base
	entity2.PropBase
}

func (entity *PropDynamicOrnament) New() entity.IEntity {
	return &PropDynamicOverride{}
}

func (entity PropDynamicOrnament) Classname() string {
	return "prop_dynamic_ornament"
}
