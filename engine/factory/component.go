package factory

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/core/interfaces"
	"github.com/galaco/Gource-Engine/engine/entity"
)

// Attaches a component to an entity, and registers it with the engine
func NewComponent(component interfaces.IComponent, owner entity.IEntity) *interfaces.IComponent {
	component.SetHandle(core.NewHandle())
	GetObjectManager().AddComponent(component, owner)
	return &component
}
