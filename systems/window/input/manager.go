package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/message/messagetype"
)

type Manager struct {
}

func (manager *Manager) Register(window *glfw.Window) {
	window.SetKeyCallback(manager.KeyCallback)
}

func (manager *Manager) Update(dt float64) {
	glfw.PollEvents()
}

func (manager *Manager) Unregister() {

}

func (manager *Manager) KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		event.GetEventManager().Dispatch(messagetype.KeyDown, &messages.KeyDown{Key: key})
	case glfw.Repeat:
		event.GetEventManager().Dispatch(messagetype.KeyHeld, &messages.KeyHeld{Key: key})
	case glfw.Release:
		event.GetEventManager().Dispatch(messagetype.KeyReleased, &messages.KeyReleased{Key: key})
	}

}