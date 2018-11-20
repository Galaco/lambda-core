package common

import (
	entity2 "github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/entity"
)

type PropDoorRotating struct {
	entity2.Base
	entity.PropBase
}

func (entity *PropDoorRotating) New() entity2.IEntity {
	return &PropDynamic{}
}

func (entity PropDoorRotating) Classname() string {
	return "prop_door_rotating"
}
