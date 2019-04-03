package common

import (
	entity2 "github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/game/entity"
)

// PropPhysics
type PropPhysics struct {
	entity2.Base
	entity.PropBase
}

// New
func (entity *PropPhysics) New() entity2.IEntity {
	return &PropPhysics{}
}

// Classname
func (entity PropPhysics) Classname() string {
	return "prop_physics"
}
