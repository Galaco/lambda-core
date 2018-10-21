package factory

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/entity"
)

// Returns a new entity, registered with the engine
func NewEntity(entity entity.IEntity) entity.IEntity {
	entity.SetHandle(core.NewHandle())

	GetObjectManager().AddEntity(entity)
	return entity
}
