package renderer

import (
	"github.com/galaco/go-me-engine/systems/renderer/gl"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/components"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
)

type Manager struct {
	base.Manager
	glContext gl.Context
}

func (manager *Manager) Register() {
	manager.glContext = gl.NewContext()
}

func (manager *Manager) Update(dt float64) {
	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)
	manager.glContext.UseProgram()

	for _,c := range factory.GetObjectManager().GetAllComponents() {
		if c.GetType() == components.T_RenderableComponent {
			resource := c.(*components.RenderableComponent).GetRenderable()
			resource.BindData()
			opengl.BindVertexArray(resource.GetVao())
			opengl.DrawArrays(opengl.TRIANGLES, 0, int32(len(resource.GetVertexData()) / 3))
		}
	}
}