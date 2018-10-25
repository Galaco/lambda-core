package renderer

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/input"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/renderer/gl"
	"github.com/galaco/Gource-Engine/engine/renderer/gl/shaders"
	"github.com/galaco/Gource-Engine/engine/renderer/gl/shaders/sky"
	"github.com/galaco/Gource-Engine/engine/scene"
	"github.com/galaco/Gource-Engine/engine/scene/world"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Manager struct {
	core.Manager
	defaultShader gl.Context
	skyShader     gl.Context

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

	// Since we only have 1 shader for now..
	manager.defaultShader.UseProgram()
	projectionUniform := manager.defaultShader.GetUniform("projection")
	projection := scene.Get().CurrentCamera().ProjectionMatrix()
	opengl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	opengl.Enable(opengl.BLEND)
	opengl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
	opengl.Enable(opengl.DEPTH_TEST)
	opengl.LineWidth(32)
	opengl.DepthFunc(opengl.LEQUAL)

	opengl.ClearColor(0, 0, 0, 1)
}

func (manager *Manager) Update(dt float64) {
	currentScene := scene.Get()
	if !currentScene.IsLoaded() {
		return
	}

	manager.defaultShader.UseProgram()
	projectionUniform := manager.defaultShader.GetUniform("projection")
	projection := scene.Get().CurrentCamera().ProjectionMatrix()
	opengl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	manager.updateRendererProperties()
	currentScene.CurrentCamera().Update(dt)

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)

	viewUniform := manager.defaultShader.GetUniform("view")
	view := scene.Get().CurrentCamera().ViewMatrix()
	opengl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	manager.drawBsp(currentScene.GetWorld())

	manager.drawStaticProps(currentScene.GetWorld().GetVisibleStaticProps())

	manager.drawSkybox(currentScene.GetWorld().GetSkybox())
}

func (manager *Manager) drawBsp(world *world.World) {
	modelUniform := manager.defaultShader.GetUniform("model")
	model := mgl32.Ident4()
	opengl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
	world.UpdateVisibilityList(scene.Get().CurrentCamera().Transform().Position)
	manager.drawModel(world.GetVisibleBsp())
}

func (manager *Manager) drawStaticProps(props []*world.StaticProp) {
	modelUniform := manager.defaultShader.GetUniform("model")

	for _,prop := range props {
		model := prop.Transform().GetTransformationMatrix()
		opengl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
		manager.drawModel(prop.GetModel())
	}
}

func (manager *Manager) drawSkybox(sky *world.Sky) {
	modelUniform := manager.defaultShader.GetUniform("model")
	model := sky.Transform().GetTransformationMatrix()
	opengl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
	manager.drawModel(sky.GetVisibleBsp())

	manager.drawStaticProps(sky.GetVisibleProps())

	manager.drawSkyMaterial(sky.GetBackdrop())
}

// render a mesh and its submeshes/primitives
func (manager *Manager) drawModel(model *model.Model) {
	for _, mesh := range model.GetMeshes() {
		// Missing materials will be flat coloured
		if mesh == nil || mesh.GetMaterial() == nil {
			// We need the fallback material
			continue
		}
		mesh.Bind()
		mesh.GetMaterial().Bind()
		if manager.renderAsWireframe == true {
			opengl.DrawArrays(opengl.LINES, 0, int32(len(mesh.Vertices()))/3)
		} else {
			opengl.DrawArrays(opengl.TRIANGLES, 0, int32(len(mesh.Vertices()))/3)
		}
	}
}

func (manager *Manager) drawSkyMaterial(skybox *model.Model) {
	var oldCullFaceMode int32
	opengl.GetIntegerv(opengl.CULL_FACE_MODE, &oldCullFaceMode)
	var oldDepthFuncMode int32
	opengl.GetIntegerv(opengl.DEPTH_FUNC, &oldDepthFuncMode)
	manager.skyShader.UseProgram()
	model := mgl32.Ident4()
	camTransform := scene.Get().CurrentCamera().Transform().Position
	model = model.Mul4(mgl32.Translate3D(camTransform.X(), camTransform.Y(), camTransform.Z()))
	model = model.Mul4(mgl32.Scale3D(20, 20, 20))
	opengl.UniformMatrix4fv(manager.skyShader.GetUniform("model"), 1, false, &model[0])
	view := scene.Get().CurrentCamera().ViewMatrix()
	opengl.UniformMatrix4fv(manager.skyShader.GetUniform("view"), 1, false, &view[0])
	projection := scene.Get().CurrentCamera().ProjectionMatrix()
	opengl.UniformMatrix4fv(manager.skyShader.GetUniform("projection"), 1, false, &projection[0])

	opengl.CullFace(opengl.FRONT)
	opengl.DepthFunc(opengl.LEQUAL)
	//DRAW
	manager.drawModel(skybox)

	opengl.CullFace(uint32(oldCullFaceMode))
	opengl.DepthFunc(uint32(oldDepthFuncMode))
}

func (manager *Manager) updateRendererProperties() {
	manager.renderAsWireframe = input.GetKeyboard().IsKeyDown(glfw.KeyX)
}
