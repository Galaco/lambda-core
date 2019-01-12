package common

import (
	"github.com/galaco/Lambda-Core/core/entity"
	entity2 "github.com/galaco/Lambda-Core/game/entity"
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
