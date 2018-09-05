package renderer

import (
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/engine/input"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/entity"
	"github.com/galaco/go-me-engine/systems/renderer/camera"
	"github.com/galaco/go-me-engine/systems/renderer/gl"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Manager struct {
	base.Manager
	glContext     gl.Context
	currentCamera camera.Camera

	renderAsWireframe bool
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
	opengl.LineWidth(32)

	opengl.ClearColor(0.5, 0.5, 0.5, 1)
}

func (manager *Manager) Update(dt float64) {
	manager.updateRendererProperties()
	manager.currentCamera.Update(dt)

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)

	modelUniform := manager.glContext.GetUniform("model")
	model := mgl32.Ident4()
	viewUniform := manager.glContext.GetUniform("view")
	view := manager.currentCamera.ViewMatrix()
	opengl.UniformMatrix4fv(viewUniform, 1, false, &view[0])


	for _, ent := range factory.GetObjectManager().GetAllEntities() {
		switch ent.(type) {
		case *entity.WorldSpawn:
			ent.(*entity.WorldSpawn).UpdateVisibilityList(manager.currentCamera.GetOwner().GetTransformComponent().Position)
			opengl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
			for _, resource := range ent.(*entity.WorldSpawn).GetPrimitives() {
				manager.drawMesh(resource)
			}
		}
	}

	for _, c := range factory.GetObjectManager().GetAllComponents() {
		switch c.(type) {
		case *components.RenderableComponent:
			modelMatrix := factory.GetObjectManager().GetEntityByHandle(c.GetOwnerHandle()).(*entity.ValveEntity).GetTransformComponent().GetTransformationMatrix()
			opengl.UniformMatrix4fv(modelUniform, 1, false, &modelMatrix[0])

			for _, resource := range c.(*components.RenderableComponent).GetRenderables() {
				manager.drawMesh(resource)
			}
		}
	}
}

// render a mesh and its submeshes/primitives
func (manager *Manager) drawMesh(resource interfaces.IGPUMesh) {
	for _, primitive := range resource.GetPrimitives() {
		// Missing materials will be flat coloured
		if primitive.GetMaterial() == nil {
			// We need the fallback material
			continue
		}
		primitive.Bind()
		primitive.GetMaterial().Bind()
		if manager.renderAsWireframe == true {
			opengl.DrawArrays(opengl.LINES, 0, int32(len(primitive.GetVertices()))/3)
		} else {
			opengl.DrawArrays(primitive.GetFaceMode(), 0, int32(len(primitive.GetVertices()))/3)
		}
	}
}

func (manager *Manager) updateRendererProperties() {
	manager.renderAsWireframe = input.GetKeyboard().IsKeyDown(glfw.KeyX)
}