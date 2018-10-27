package common

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	entity2 "github.com/galaco/Gource-Engine/entity"
)

type PropPhysicsMultiplayer struct {
	entity.Base
	entity2.PropBase
}

func (entity *PropPhysicsMultiplayer) New() entity.IEntity {
	return &PropPhysicsMultiplayer{}
}

func (entity PropPhysicsMultiplayer) Classname() string {
	return "prop_physics_multiplayer"
}