package window

import (
	"github.com/galaco/Gource/engine/base"
	"github.com/galaco/Gource/systems/window/input"
	"github.com/galaco/Gource/systems/window/window"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Manager struct {
	base.Manager
	window *glfw.Window
	input  input.Manager
}

func (manager *Manager) Register() {
	manager.window = window.Create(640, 480, "test_window")
	manager.input.Register(manager.window)
}

func (manager *Manager) Update(dt float64) {
	manager.input.Update(0)
}

func (manager *Manager) Unregister() {
	manager.input.Unregister()
	glfw.Terminate()
}

func (manager *Manager) PostUpdate() {
	manager.window.SwapBuffers()
}
