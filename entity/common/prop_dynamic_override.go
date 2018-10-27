package common

import "github.com/galaco/Gource-Engine/engine/entity"

type PropDynamicOverride struct {
	PropDynamic
}

func (entity *PropDynamicOverride) New() entity.IEntity {
	return &PropDynamicOverride{}
}

func (entity PropDynamicOverride) Classname() string {
	return "prop_dynamic_override"
}