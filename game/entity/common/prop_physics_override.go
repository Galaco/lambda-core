package common

import (
	"github.com/galaco/Lambda-Core/core/entity"
	entity2 "github.com/galaco/Lambda-Core/game/entity"
)

// PropPhysicsOverride
type PropPhysicsOverride struct {
	entity.Base
	entity2.PropBase
}

//New
func (entity *PropPhysicsOverride) New() entity.IEntity {
	return &PropPhysicsOverride{}
}

// Classname
func (entity PropPhysicsOverride) Classname() string {
	return "prop_physics_override"
}
