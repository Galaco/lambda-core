package entity

import (
	"github.com/galaco/Gource/engine/base"
	"github.com/galaco/source-tools-common/entity"
)

type ValveEntity struct {
	base.Entity
	Definition *entity.Entity
}

func NewEntity(entityDefinition *entity.Entity) *ValveEntity {
	return &ValveEntity{
		Definition: entityDefinition,
	}
}
