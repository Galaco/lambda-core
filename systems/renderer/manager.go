package renderer

import "github.com/galaco/bsp-viewer/systems/renderer/gl"

type Manager struct {
	glContext gl.Context
}

func (manager *Manager) Register() {
	manager.glContext = gl.NewContext()

}

func (manager *Manager) Update(dt float64) {

}

func (manager *Manager) Unregister() {

}