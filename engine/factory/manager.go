package factory

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/core"
)

type Manager struct {
	entities []interfaces.IEntity
	components []interfaces.IComponent
}

func (manager *Manager) GetAllEntities() []interfaces.IEntity {
	return manager.entities
}

func (manager *Manager) GetAllComponents() []interfaces.IComponent {
	return manager.components
}

func (manager *Manager) GetEntityByHandle(handle core.Handle) interfaces.IEntity {
	if len(manager.entities) < int(handle) {
		return nil
	}
	return manager.entities[handle]
}

func (manager *Manager) GetComponentByHandle(handle core.Handle) interfaces.IComponent {
	if len(manager.components) < int(handle) {
		return nil
	}
	return manager.components[handle]
}

func (manager *Manager) AddEntity(ent interfaces.IEntity) {
	manager.entities = append(manager.entities, ent)
}

func (manager *Manager) AddComponent(component interfaces.IComponent, ent interfaces.IEntity) {
	component.SetOwnerHandle(ent.GetHandle())
	manager.components = append(manager.components, component)
	ent.AddComponent(component.GetHandle())
	component.Initialize()
}



// There can be only 1 instance.
// I guess its a gross singleton?
// Would rather it be properly static, but golang no likey :(
var objectManager Manager

func GetObjectManager() *Manager{
	return &objectManager
}