package gl

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/renderer/gl/shaders"
	"github.com/galaco/Gource-Engine/engine/renderer/gl/shaders/sky"
	"github.com/galaco/Gource-Engine/engine/scene"
	"github.com/galaco/Gource-Engine/engine/scene/world"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//OpenGL renderer
type Renderer struct {
	defaultShader Context
	skyShader     Context

	uniformMap map[string]int32

	vertexDrawMode uint32
}

// Preparation function
// Loads shaders and sets necessary constants for opengls state machine
func (manager *Renderer) LoadShaders() {
	manager.defaultShader = NewContext()
	manager.defaultShader.AddShader(shaders.Vertex, opengl.VERTEX_SHADER)
	manager.defaultShader.AddShader(shaders.Fragment, opengl.FRAGMENT_SHADER)
	manager.defaultShader.Finalize()
	manager.skyShader = NewContext()
	manager.skyShader.AddShader(sky.Vertex, opengl.VERTEX_SHADER)
	manager.skyShader.AddShader(sky.Fragment, opengl.FRAGMENT_SHADER)
	manager.skyShader.Finalize()

	manager.uniformMap = map[string]int32{}

	opengl.Enable(opengl.BLEND)
	opengl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
	opengl.Enable(opengl.DEPTH_TEST)
	opengl.LineWidth(32)
	opengl.DepthFunc(opengl.LEQUAL)

	opengl.ClearColor(0, 0, 0, 1)
}

// Called at the start of a frame
func (manager *Renderer) StartFrame(camera *entity.Camera) {
	manager.defaultShader.UseProgram()
	manager.uniformMap["model"] = manager.defaultShader.GetUniform("model")
	manager.uniformMap["projection"] = manager.defaultShader.GetUniform("projection")
	projection := camera.ProjectionMatrix()
	opengl.UniformMatrix4fv(manager.uniformMap["projection"], 1, false, &projection[0])

	viewUniform := manager.defaultShader.GetUniform("view")
	view := camera.ViewMatrix()
	opengl.UniformMatrix4fv(viewUniform, 1, false, &view[0])

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)
}

// Called at the end of a frame
func (manager *Renderer) EndFrame() {

}

// Draw the main bsp world
func (manager *Renderer) DrawBsp(world *world.World) {
	manager.DrawModel(world.VisibleWorld().Bsp(), mgl32.Ident4())
}

// Draw passed static props
func (manager *Renderer) DrawStaticProps(props []*world.StaticProp) {
	for _, prop := range props {
		manager.DrawModel(prop.GetModel(), prop.Transform().GetTransformationMatrix())
	}
}

// Draw skybox (bsp model, staticprops, sky material)
func (manager *Renderer) DrawSkybox(sky *world.Sky) {
	if sky == nil {
		return
	}

	if sky.GetVisibleBsp() != nil {
		manager.DrawModel(sky.GetVisibleBsp(), sky.Transform().GetTransformationMatrix())
	}

	manager.DrawStaticProps(sky.GetVisibleProps())

	manager.DrawSkyMaterial(sky.GetBackdrop())
}

// Render a mesh and its submeshes/primitives
func (manager *Renderer) DrawModel(model *model.Model, transform mgl32.Mat4) {
	opengl.UniformMatrix4fv(manager.uniformMap["model"], 1, false, &transform[0])

	for _, mesh := range model.GetMeshes() {
		// Missing materials will be flat coloured
		if mesh == nil || mesh.GetMaterial() == nil {
			// We need the fallback material
			continue
		}
		mesh.Bind()
		mesh.GetMaterial().Bind()
		opengl.DrawArrays(manager.vertexDrawMode, 0, int32(len(mesh.Vertices()))/3)
	}
}

// Render the sky material
func (manager *Renderer) DrawSkyMaterial(skybox *model.Model) {
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
	manager.DrawModel(skybox, mgl32.Ident4())

	opengl.CullFace(uint32(oldCullFaceMode))
	opengl.DepthFunc(uint32(oldDepthFuncMode))
}

// Change the draw format.
func (manager *Renderer) SetWireframeMode(mode bool) {
	if mode == true {
		manager.vertexDrawMode = opengl.LINES
	} else {
		manager.vertexDrawMode = opengl.TRIANGLES
	}
}

func NewRenderer() *Renderer {
	r := Renderer{}
	r.SetWireframeMode(false)

	return &r
}
