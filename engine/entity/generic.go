package entity

import (
	"github.com/galaco/Gource-Engine/engine/core"
	entity2 "github.com/galaco/source-tools-common/entity"
)

// GenericEntity can be used to substitute out an entity
// that doesn't have a known implementation
type GenericEntity struct {
	Base
}

// NewGenericEntity returns new Entity
func NewGenericEntity(definition *entity2.Entity) *GenericEntity {
	ent := &GenericEntity{
		Base: Base{
			keyValues: definition,
			handle:    core.NewHandle(),
		},
	}

	return ent
}
