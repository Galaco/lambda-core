package event

import (
	"github.com/galaco/bsp-viewer/engine/core"
	"github.com/galaco/bsp-viewer/engine/interfaces"
)

type Manager struct {
	listenerMap map[string]map[core.Handle]interfaces.IComponent
}

//Register a new component to listen to an event
func (manager *Manager) RegisterEvent(eventName string, component interfaces.IComponent) core.Handle{
	handle := core.NewHandle()
	if _,ok := manager.listenerMap[eventName]; !ok {
		manager.listenerMap[eventName] = make(map[core.Handle]interfaces.IComponent)
	}
	manager.listenerMap[eventName][handle] = component

	return handle
}

//Remove a listener from listening for an event
func (manager *Manager) UnregisterEvent(eventName string, handle core.Handle) {
	if _, ok := manager.listenerMap[eventName][handle]; ok {
		delete(manager.listenerMap[eventName], handle)
	}
}

//Fires an event to all listening components
func (manager *Manager) FireEvent(eventName string, message interfaces.IMessage) {
	message.SetType(eventName)
	for _,component := range manager.listenerMap[eventName] {
		component.ReceiveMessage(message)
	}
}







var eventManager Manager

func GetEventManager() *Manager {
	if eventManager.listenerMap == nil {
		eventManager.listenerMap = make(map[string]map[core.Handle]interfaces.IComponent)
	}
	return &eventManager
}