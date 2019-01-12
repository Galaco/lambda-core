package common

import (
	entity2 "github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/game/entity"
)

type PropPhysics struct {
	entity2.Base
	entity.PropBase
}

func (entity *PropPhysics) New() entity2.IEntity {
	return &PropPhysics{}
}

func (entity PropPhysics) Classname() string {
	return "prop_physics"
}
