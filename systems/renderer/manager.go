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
	"github.com/galaco/go-me-engine/systems/renderer/gl/shaders"
	"github.com/galaco/go-me-engine/systems/renderer/gl/shaders/sky"
)

type Manager struct {
	base.Manager
	defaultShader     gl.Context
	skyShader     gl.Context
	currentCamera camera.Camera

	renderAsWireframe bool
}

func (manager *Manager) Register() {
	manager.defaultShader = gl.NewContext()
	manager.defaultShader.AddShader(shaders.Vertex, opengl.VERTEX_SHADER)
	manager.defaultShader.AddShader(shaders.Fragment, opengl.FRAGMENT_SHADER)
	manager.defaultShader.Finalize()
	manager.skyShader = gl.NewContext()
	manager.skyShader.AddShader(sky.Vertex, opengl.VERTEX_SHADER)
	manager.skyShader.AddShader(sky.Fragment, opengl.FRAGMENT_SHADER)
	manager.skyShader.Finalize()


	manager.currentCamera.Initialize()

	// Since we only have 1 shader for now..
	manager.defaultShader.UseProgram()
	projectionUniform := manager.defaultShader.GetUniform("projection")
	projection := manager.currentCamera.ProjectionMatrix()
	opengl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	opengl.Enable(opengl.BLEND)
	opengl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
	opengl.Enable(opengl.DEPTH_TEST)
	opengl.LineWidth(32)
	opengl.DepthFunc(opengl.LEQUAL)

	opengl.ClearColor(0.5, 0.5, 0.5, 1)
}

func (manager *Manager) Update(dt float64) {
	manager.updateRendererProperties()
	manager.currentCamera.Update(dt)

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)

	modelUniform := manager.defaultShader.GetUniform("model")
	model := mgl32.Ident4()
	viewUniform := manager.defaultShader.GetUniform("view")
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
		case *components.Skybox:
			manager.drawSky(c.(*components.Skybox))
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

func (manager *Manager) drawSky(skybox *components.Skybox) {
	var oldCullFaceMode int32
	opengl.GetIntegerv(opengl.CULL_FACE_MODE, &oldCullFaceMode)
	var oldDepthFuncMode int32
	opengl.GetIntegerv(opengl.DEPTH_FUNC, &oldDepthFuncMode)
	manager.skyShader.UseProgram()
	model := mgl32.Ident4()
	camTransform := manager.currentCamera.GetOwner().GetTransformComponent().Position
	model = model.Mul4(mgl32.Translate3D(camTransform.X(), camTransform.Y(), camTransform.Z()))
	model = model.Mul4(mgl32.Scale3D(20, 20, 20))
	opengl.UniformMatrix4fv(manager.skyShader.GetUniform("model"), 1, false, &model[0])
	view := manager.currentCamera.ViewMatrix()
	opengl.UniformMatrix4fv(manager.skyShader.GetUniform("view"), 1, false, &view[0])
	projection := manager.currentCamera.ProjectionMatrix()
	opengl.UniformMatrix4fv(manager.skyShader.GetUniform("projection"), 1, false, &projection[0])

	opengl.CullFace(opengl.FRONT)
	opengl.DepthFunc(opengl.LEQUAL)
	//DRAW
	manager.drawMesh(skybox.GetRenderables()[0])

	// Set back to default shader.
	// Why? Only called 1 time per frame
	manager.defaultShader.UseProgram()
	opengl.UniformMatrix4fv(manager.defaultShader.GetUniform("view"), 1, false, &view[0])
	opengl.UniformMatrix4fv(manager.defaultShader.GetUniform("projection"), 1, false, &projection[0])


	opengl.CullFace(uint32(oldCullFaceMode))
	opengl.DepthFunc(uint32(oldDepthFuncMode))
}

func (manager *Manager) updateRendererProperties() {
	manager.renderAsWireframe = input.GetKeyboard().IsKeyDown(glfw.KeyX)
}