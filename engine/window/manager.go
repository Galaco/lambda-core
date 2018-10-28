package window

import (
	"github.com/galaco/Gource-Engine/engine/config"
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/window/input"
	"github.com/galaco/Gource-Engine/engine/window/window"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Manager struct {
	core.Manager
	window *glfw.Window
	input  input.Manager
}

func (manager *Manager) Register() {
	manager.window = window.Create(config.Get().Video.Width, config.Get().Video.Height, "test_window")
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
