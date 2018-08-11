package factory

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/core"
)

func NewComponent(component interfaces.IComponent, owner *interfaces.IEntity) *interfaces.IComponent{
	component.SetHandle(core.NewHandle())
	GetObjectManager().AddComponent(component, *owner)
	return &component
}
