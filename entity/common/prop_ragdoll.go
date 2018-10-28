package common

import (
	entity2 "github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/entity"
)

type PropRagdoll struct {
	entity2.Base
	entity.PropBase
}

func (entity *PropRagdoll) New() entity2.IEntity {
	return &PropRagdoll{}
}

func (entity PropRagdoll) Classname() string {
	return "prop_ragdoll"
}
