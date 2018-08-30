package factory

import (
	"github.com/galaco/go-me-engine/engine/core"
	"github.com/galaco/go-me-engine/engine/interfaces"
)

// Returns a new entity, registered with the engine
func NewEntity(entity interfaces.IEntity) *interfaces.IEntity {
	entity.SetHandle(core.NewHandle())

	GetObjectManager().AddEntity(entity)
	return &entity
}
