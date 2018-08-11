package factory

import (
	"github.com/galaco/bsp-viewer/engine/interfaces"
	"github.com/galaco/bsp-viewer/engine/core"
)

func NewEntity(entity interfaces.IEntity) *interfaces.IEntity{
	entity.SetHandle(core.NewHandle())

	GetObjectManager().AddEntity(entity)
	return &entity
}