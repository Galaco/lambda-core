package factory

import (
	"github.com/galaco/Gource-Engine/engine/component"
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/entity"
)

// Attaches a component to an entity, and registers it with the engine
func NewComponent(component component.IComponent, owner entity.IEntity) *component.IComponent {
	component.SetHandle(core.NewHandle())
	GetObjectManager().AddComponent(component, owner)
	return &component
}
