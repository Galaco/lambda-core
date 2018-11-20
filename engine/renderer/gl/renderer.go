package gl

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/renderer/gl/shaders"
	"github.com/galaco/Gource-Engine/engine/renderer/gl/shaders/sky"
	"github.com/galaco/Gource-Engine/engine/scene"
	"github.com/galaco/Gource-Engine/engine/scene/world"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"log"
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

	opengl.Enable(opengl.BLEND)
	opengl.BlendFunc(opengl.SRC_ALPHA, opengl.ONE_MINUS_SRC_ALPHA)
	opengl.Enable(opengl.DEPTH_TEST)
	opengl.LineWidth(32)
	opengl.DepthFunc(opengl.LEQUAL)
	opengl.Enable(opengl.CULL_FACE)
	opengl.CullFace(opengl.BACK)
	opengl.FrontFace(opengl.CW)

	opengl.ClearColor(0, 0, 0, 1)
}

var numCalls = 0

// Called at the start of a frame
func (manager *Renderer) StartFrame(camera *entity.Camera) {
	manager.defaultShader.UseProgram()

	//matrixes
	manager.uniformMap["model"] = manager.defaultShader.GetUniform("model")
	manager.uniformMap["projection"] = manager.defaultShader.GetUniform("projection")
	projection := camera.ProjectionMatrix()
	opengl.UniformMatrix4fv(manager.uniformMap["projection"], 1, false, &projection[0])
	manager.uniformMap["view"] = manager.defaultShader.GetUniform("view")
	view := camera.ViewMatrix()
	opengl.UniformMatrix4fv(manager.uniformMap["view"], 1, false, &view[0])

	//material properties
	manager.uniformMap["baseTextureSampler"] = manager.defaultShader.GetUniform("baseTextureSampler")

	manager.uniformMap["useLightmap"] = manager.defaultShader.GetUniform("useLightmap")
	manager.uniformMap["lightmapTextureSampler"] = manager.defaultShader.GetUniform("lightmapTextureSampler")
	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)
}

// Called at the end of a frame
func (manager *Renderer) EndFrame() {
	//if glError := opengl.GetError(); glError != opengl.NO_ERROR {
	//	debug.Error("error: %d\n", glError)
	//}
	//debug.Logf("Calls: %d", numCalls)
	numCalls = 0
}

// Draw the main bsp world
func (manager *Renderer) DrawBsp(world *world.World) {
	modelMatrix := mgl32.Ident4()
	opengl.UniformMatrix4fv(manager.uniformMap["model"], 1, false, &modelMatrix[0])
	manager.BindMesh(world.Bsp().Mesh())
	log.Println(len(world.VisibleClusters()))
	for _,cluster := range world.VisibleClusters() {
		for _,face := range cluster.Faces {
			manager.DrawFace(&face)
		}
	}
	for _,cluster := range world.VisibleClusters() {
		for _,prop := range cluster.StaticProps {
			manager.DrawModel(prop.GetModel(), prop.Transform().GetTransformationMatrix())
		}
	}
}

// Draw skybox (bsp model, staticprops, sky material)
func (manager *Renderer) DrawSkybox(sky *world.Sky) {
	if sky == nil {
		return
	}

	if sky.GetVisibleBsp() != nil {
		modelMatrix := sky.Transform().GetTransformationMatrix()
		opengl.UniformMatrix4fv(manager.uniformMap["model"], 1, false, &modelMatrix[0])
		manager.BindMesh(sky.GetVisibleBsp().Mesh())
		for _,cluster := range sky.GetClusterLeafs() {
			for _,face := range cluster.Faces {
				manager.DrawFace(&face)
			}
		}
		for _,cluster := range sky.GetClusterLeafs() {
			for _,prop := range cluster.StaticProps {
				manager.DrawModel(prop.GetModel(), prop.Transform().GetTransformationMatrix())
			}
		}
	}

	//manager.DrawSkyMaterial(sky.GetBackdrop())
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
		manager.BindMesh(mesh)
		opengl.DrawArrays(manager.vertexDrawMode, 0, int32(len(mesh.Vertices()))/3)

		numCalls++
	}
}

func (manager *Renderer) BindMesh(target mesh.IMesh) {
	target.Bind()
	// $basetexture
	if target.GetMaterial() != nil {
		target.GetMaterial().Bind()
		opengl.Uniform1i(manager.uniformMap["baseTextureSampler"], 0)
	}
	// Bind lightmap texture if it exists
	if target.GetLightmap() != nil {
		opengl.Uniform1i(manager.uniformMap["useLightmap"], 1)
		opengl.Uniform1i(manager.uniformMap["lightmapTextureSampler"], 1)
		target.GetLightmap().Bind()
	} else {
		opengl.Uniform1i(manager.uniformMap["useLightmap"], 0)
	}
}

func (manager *Renderer) DrawFace(target *mesh.Face) {
	// Skip materialless faces
	if target.Material() == nil {
		return
	}
	// $basetexture
	target.Material().Bind()
	opengl.Uniform1i(manager.uniformMap["baseTextureSampler"], 0)

	// Bind lightmap texture if it exists
	if target.IsLightmapped() == true {
		opengl.Uniform1i(manager.uniformMap["useLightmap"], 1)
		opengl.Uniform1i(manager.uniformMap["lightmapTextureSampler"], 1)
		target.Lightmap().Bind()
	} else {
		opengl.Uniform1i(manager.uniformMap["useLightmap"], 0)
	}
	opengl.DrawArrays(manager.vertexDrawMode, target.Offset(), target.Length())
}

// Render the sky material
func (manager *Renderer) DrawSkyMaterial(skybox *model.Model) {
	if skybox == nil {
		return
	}
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
	r.uniformMap = map[string]int32{}

	return &r
}
