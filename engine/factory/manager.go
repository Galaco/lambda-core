package factory

import (
	"github.com/galaco/Gource-Engine/engine/component"
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/entity"
)

// Game object manager
// Store entities and components
type Manager struct {
	entities   map[core.Handle]entity.IEntity
	components map[core.Handle]component.IComponent
}

// Returns all existing entities
func (manager *Manager) GetAllEntities() map[core.Handle]entity.IEntity {
	return manager.entities
}

// Returns all existing components
func (manager *Manager) GetAllComponents() map[core.Handle]component.IComponent {
	return manager.components
}

// Find a specific entity by its unique name
func (manager *Manager) GetEntityByHandle(handle core.Handle) entity.IEntity {
	return manager.entities[handle]
}

// Find a specific component by its unique name
func (manager *Manager) GetComponentByHandle(handle core.Handle) component.IComponent {
	return manager.components[handle]
}

// Add a new entity
func (manager *Manager) AddEntity(ent entity.IEntity) {
	manager.entities[ent.GetHandle()] = ent
}

// Add a new component, registered against an existing entity
func (manager *Manager) AddComponent(component component.IComponent, ent entity.IEntity) {
	component.SetOwnerHandle(ent.GetHandle())
	manager.components[component.GetHandle()] = component
	ent.AddComponent(component.GetHandle())
	component.Initialize()
}

// There can be only 1 instance.
// I guess its a gross singleton?
// Would rather it be properly static, but golang no likey :(
var objectManager Manager

func GetObjectManager() *Manager {
	if objectManager.components == nil {
		objectManager.entities = make(map[core.Handle]entity.IEntity)
		objectManager.components = make(map[core.Handle]component.IComponent)
	}
	return &objectManager
}
