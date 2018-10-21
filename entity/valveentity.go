package entity

import (
	"github.com/galaco/source-tools-common/entity"
)

type ValveEntity struct {
	entity.Entity
	Definition *entity.Entity
}

func NewEntity(entityDefinition *entity.Entity) *ValveEntity {
	return &ValveEntity{
		Definition: entityDefinition,
	}
}
