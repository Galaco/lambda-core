package factory

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/interfaces"
)

// Returns a new entity, registered with the engine
func NewEntity(entity interfaces.IEntity) interfaces.IEntity {
	entity.SetHandle(core.NewHandle())

	GetObjectManager().AddEntity(entity)
	return entity
}
