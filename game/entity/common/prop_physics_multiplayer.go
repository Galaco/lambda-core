package common

import (
	"github.com/galaco/Lambda-Core/core/entity"
	entity2 "github.com/galaco/Lambda-Core/game/entity"
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
