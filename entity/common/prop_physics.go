package common

import (
	entity2 "github.com/galaco/Gource-Engine/engine/entity"
)

type PropPhysics struct {
	entity2.Base
}

func (entity *PropPhysics) New() entity2.IEntity {
	return &PropPhysics{}
}

func (entity PropPhysics) Classname() string {
	return "prop_physics"
}
