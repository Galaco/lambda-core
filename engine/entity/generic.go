package entity

import (
	"github.com/galaco/Gource-Engine/engine/core"
	entity2 "github.com/galaco/source-tools-common/entity"
)

type GenericEntity struct {
	*Base
}

func NewGenericEntity(definition *entity2.Entity) GenericEntity {
	ent := GenericEntity{
		Base: &Base{
			keyValues: definition,
			handle:    core.NewHandle(),
		},
	}

	return ent
}
