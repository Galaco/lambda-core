package factory

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/core"
)

type Manager struct {
	entities map[core.Handle]interfaces.IEntity
	components map[core.Handle]interfaces.IComponent
}

func (manager *Manager) GetAllEntities() map[core.Handle]interfaces.IEntity {
	return manager.entities
}

func (manager *Manager) GetAllComponents() map[core.Handle]interfaces.IComponent {
	return manager.components
}

func (manager *Manager) GetEntityByHandle(handle core.Handle) interfaces.IEntity {
	return manager.entities[handle]
}

func (manager *Manager) GetComponentByHandle(handle core.Handle) interfaces.IComponent {
	return manager.components[handle]
}

func (manager *Manager) AddEntity(ent interfaces.IEntity) {
	manager.entities[ent.GetHandle()] = ent
}

func (manager *Manager) AddComponent(component interfaces.IComponent, ent interfaces.IEntity) {
	component.SetOwnerHandle(ent.GetHandle())
	manager.components[component.GetHandle()] = component
	ent.AddComponent(component.GetHandle())
	component.Initialize()
}



// There can be only 1 instance.
// I guess its a gross singleton?
// Would rather it be properly static, but golang no likey :(
var objectManager Manager

func GetObjectManager() *Manager{
	if objectManager.components == nil {
		objectManager.entities = make(map[core.Handle]interfaces.IEntity)
		objectManager.components = make(map[core.Handle]interfaces.IComponent)
	}
	return &objectManager
}