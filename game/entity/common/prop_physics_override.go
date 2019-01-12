package common

import (
	"github.com/galaco/Lambda-Core/core/entity"
	entity2 "github.com/galaco/Lambda-Core/game/entity"
)

type PropPhysicsOverride struct {
	entity.Base
	entity2.PropBase
}

func (entity *PropPhysicsOverride) New() entity.IEntity {
	return &PropPhysicsOverride{}
}

func (entity PropPhysicsOverride) Classname() string {
	return "prop_physics_override"
}
