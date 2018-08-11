package factory

import (
	"github.com/galaco/bsp-viewer/engine/interfaces"
	"github.com/galaco/bsp-viewer/engine/core"
)

func NewComponent(component interfaces.IComponent, owner *interfaces.IEntity) *interfaces.IComponent{
	component.SetHandle(core.NewHandle())
	GetObjectManager().AddComponent(component, *owner)
	return &component
}
