package event

import (
	"github.com/galaco/go-me-engine/engine/core"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"sync"
)

type Manager struct {
	listenerMap map[core.EventId]map[core.Handle]interfaces.IEventListenable
	mu sync.Mutex
	eventQueue []*QueueItem
	runAsync bool
}

//Register a new component to listen to an event
func (manager *Manager) Listen(eventName core.EventId, component interfaces.IEventListenable) core.Handle{
	handle := core.NewHandle()
	manager.mu.Lock()
	if _,ok := manager.listenerMap[eventName]; !ok {
		manager.listenerMap[eventName] = make(map[core.Handle]interfaces.IEventListenable)
	}
	manager.listenerMap[eventName][handle] = component
	manager.mu.Unlock()

	return handle
}

// Runs the event queue in its own go routine
func (manager *Manager) RunConcurrent() {
	// Block double-running
	if manager.runAsync == true {
		return
	}
	manager.runAsync = true
	go func() {
		for manager.runAsync == true {
			manager.mu.Lock()
			queue := manager.eventQueue
			manager.mu.Unlock()

			if len(queue) > 0 {
				// FIFO - ensure dispatch order, and concurrency integrity
				item := queue[0]
				manager.mu.Lock()
				manager.eventQueue = manager.eventQueue[1:]

				// Fire event
				listeners := manager.listenerMap[item.EventName]
				manager.mu.Unlock()
				for _,component := range listeners {
					component.ReceiveMessage(item.Message)
				}
			}
		}
	}()
}

//Remove a listener from listening for an event
func (manager *Manager) Unlisten(eventName core.EventId, handle core.Handle) {
	manager.mu.Lock()
	if _, ok := manager.listenerMap[eventName][handle]; ok {
		delete(manager.listenerMap[eventName], handle)
	}
	manager.mu.Unlock()
}

//Fires an event to all listening components
func (manager *Manager) Dispatch(eventName core.EventId, message interfaces.IMessage) {
	message.SetType(eventName)
	queueItem := &QueueItem{
		EventName: eventName,
		Message: message,
	}
	manager.mu.Lock()
	manager.eventQueue = append(manager.eventQueue, queueItem)
	manager.mu.Unlock()
}

func (manager *Manager) Unregister() {
	// Ensure async event queue is halted
	manager.runAsync = false
}







var eventManager Manager

func GetEventManager() *Manager {
	if eventManager.listenerMap == nil {
		eventManager.listenerMap = make(map[core.EventId]map[core.Handle]interfaces.IEventListenable)
	}
	return &eventManager
}