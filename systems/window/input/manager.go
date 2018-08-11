package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/galaco/bsp-viewer/engine/event"
	"github.com/galaco/bsp-viewer/message/input"
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
		event.GetEventManager().FireEvent(event.KeyDown, &input.KeyDown{Key: key})
	case glfw.Repeat:
		event.GetEventManager().FireEvent(event.KeyHeld, &input.KeyHeld{Key: key})
	case glfw.Release:
		event.GetEventManager().FireEvent(event.KeyReleased, &input.KeyReleased{Key: key})
	}

}