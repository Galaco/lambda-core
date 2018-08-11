package factory

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/core"
)

func NewEntity(entity interfaces.IEntity) *interfaces.IEntity{
	entity.SetHandle(core.NewHandle())

	GetObjectManager().AddEntity(entity)
	return &entity
}