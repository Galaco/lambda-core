package renderer

import (
	"github.com/galaco/go-me-engine/systems/renderer/gl"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/components"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/galaco/go-me-engine/systems/renderer/camera"
)

type Manager struct {
	base.Manager
	glContext gl.Context
	currentCamera camera.Camera
}

func (manager *Manager) Register() {
	manager.glContext = gl.NewContext()
	manager.currentCamera.Initialize()

	// Since we only have 1 shader for now..
	manager.glContext.UseProgram()
	projectionUniform := manager.glContext.GetUniform("projection")
	projection := manager.currentCamera.ProjectionMatrix()
	opengl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	opengl.Enable(opengl.BLEND)
	opengl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
}

func (manager *Manager) Update(dt float64) {
	manager.currentCamera.Update(dt)

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)
//	manager.glContext.UseProgram()

	modelUniform := manager.glContext.GetUniform("model")
	model := manager.currentCamera.ModelMatrix()
	opengl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
	viewUniform := manager.glContext.GetUniform("view")
	view := manager.currentCamera.ViewMatrix()
	opengl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	for _,c := range factory.GetObjectManager().GetAllComponents() {
		if c.GetType() == components.T_RenderableComponent {
			for _,resource := range c.(*components.RenderableComponent).GetRenderables() {
				resource.Bind()
				for _, primitive := range resource.GetPrimitives() {
					// For now, just skip faces with no material
					if primitive.GetMaterial() == nil {
						continue
					}
					primitive.Bind()
					primitive.GetMaterial().Bind()
					opengl.DrawElements(primitive.GetFaceMode(), int32(len(primitive.GetIndices())), opengl.UNSIGNED_SHORT, opengl.Ptr(primitive.GetIndices()))
				}
			}
		}
	}
}