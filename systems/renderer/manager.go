package renderer

import "github.com/galaco/go-me-engine/systems/renderer/gl"

type Manager struct {
	glContext gl.Context
}

func (manager *Manager) Register() {
	manager.glContext = gl.NewContext()

}

func (manager *Manager)  RunConcurrent() {

}

func (manager *Manager) Update(dt float64) {

}

func (manager *Manager) Unregister() {

}