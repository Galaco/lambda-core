package common

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	entity2 "github.com/galaco/Gource-Engine/entity"
)

type PropDynamicOverride struct {
	entity.Base
	entity2.PropBase
}

func (entity *PropDynamicOverride) New() entity.IEntity {
	return &PropDynamicOverride{}
}

func (entity PropDynamicOverride) Classname() string {
	return "prop_dynamic_override"
}
