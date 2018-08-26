package renderer

import (
	"github.com/galaco/go-me-engine/systems/renderer/gl"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/components"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/galaco/go-me-engine/systems/renderer/camera"
	"github.com/galaco/go-me-engine/engine/interfaces"
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
	opengl.Enable(opengl.DEPTH_TEST)
}

func (manager *Manager) Update(dt float64) {
	manager.currentCamera.Update(dt)

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)

	modelUniform := manager.glContext.GetUniform("model")
	model := manager.currentCamera.ModelMatrix()
	opengl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
	viewUniform := manager.glContext.GetUniform("view")
	view := manager.currentCamera.ViewMatrix()
	opengl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	for _,c := range factory.GetObjectManager().GetAllComponents() {
		switch c.GetType() {
		//case components.T_RenderableComponent:
		//	for _,resource := range c.(*components.RenderableComponent).GetRenderables() {
		//		manager.drawMesh(resource)
		//	}
		case components.T_BspComponent:
			c.(*components.BspComponent).UpdateVisibilityList(manager.currentCamera.GetOwner().GetTransformComponent().Position)
			for _,resource := range c.(*components.BspComponent).GetRenderables() {
				manager.drawMesh(resource)
			}
		}
	}
}

func (manager *Manager) drawMesh(resource interfaces.IGPUMesh) {
	for _, primitive := range resource.GetPrimitives() {
		// Missing materials will be flat coloured
		if primitive.GetMaterial() == nil {
			// We need the fall backmaterial
			continue
		}
		primitive.Bind()
		primitive.GetMaterial().Bind()
		opengl.DrawArrays(primitive.GetFaceMode(), 0, int32(len(primitive.GetVertices())) / 3)
	}
}