package common

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	entity2 "github.com/galaco/Gource-Engine/entity"
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
