package gl

import (
	"github.com/galaco/Gource-Engine/client/renderer/gl/shaders"
	"github.com/galaco/Gource-Engine/client/renderer/gl/shaders/sky"
	"github.com/galaco/Gource-Engine/client/scene/world"
	"github.com/galaco/Gource-Engine/core/entity"
	"github.com/galaco/Gource-Engine/core/mesh"
	"github.com/galaco/Gource-Engine/core/model"
	opengl "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//OpenGL renderer
type Renderer struct {
	lightmappedGenericShader Context
	skyShader                Context

	currentShaderId uint32

	uniformMap map[uint32]map[string]int32

	vertexDrawMode uint32

	matrixes struct {
		view       mgl32.Mat4
		projection mgl32.Mat4
	}
}

// Preparation function
// Loads shaders and sets necessary constants for opengls state machine
func (manager *Renderer) LoadShaders() {
	manager.lightmappedGenericShader = NewContext()
	manager.lightmappedGenericShader.AddShader(shaders.Vertex, opengl.VERTEX_SHADER)
	manager.lightmappedGenericShader.AddShader(shaders.Fragment, opengl.FRAGMENT_SHADER)
	manager.lightmappedGenericShader.Finalize()
	manager.skyShader = NewContext()
	manager.skyShader.AddShader(sky.Vertex, opengl.VERTEX_SHADER)
	manager.skyShader.AddShader(sky.Fragment, opengl.FRAGMENT_SHADER)
	manager.skyShader.Finalize()

	//matrixes
	skyShaderMap := map[string]int32{}
	skyShaderMap["model"] = manager.skyShader.GetUniform("model")
	skyShaderMap["projection"] = manager.skyShader.GetUniform("projection")
	skyShaderMap["view"] = manager.skyShader.GetUniform("view")
	skyShaderMap["cubemapTexture"] = manager.lightmappedGenericShader.GetUniform("cubemapTexture")
	manager.uniformMap[manager.skyShader.Id()] = skyShaderMap

	manager.lightmappedGenericShader.UseProgram()
	lightmappedGenericShaderMap := map[string]int32{}
	lightmappedGenericShaderMap["model"] = manager.lightmappedGenericShader.GetUniform("model")
	lightmappedGenericShaderMap["projection"] = manager.lightmappedGenericShader.GetUniform("projection")
	lightmappedGenericShaderMap["view"] = manager.lightmappedGenericShader.GetUniform("view")
	//material properties
	lightmappedGenericShaderMap["baseTextureSampler"] = manager.lightmappedGenericShader.GetUniform("baseTextureSampler")
	lightmappedGenericShaderMap["useLightmap"] = manager.lightmappedGenericShader.GetUniform("useLightmap")
	lightmappedGenericShaderMap["lightmapTextureSampler"] = manager.lightmappedGenericShader.GetUniform("lightmapTextureSampler")
	manager.uniformMap[manager.lightmappedGenericShader.Id()] = lightmappedGenericShaderMap

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
	manager.matrixes.projection = camera.ProjectionMatrix()
	manager.matrixes.view = camera.ViewMatrix()

	// Sky
	manager.skyShader.UseProgram()
	manager.setShader(manager.skyShader.Id())
	opengl.UniformMatrix4fv(manager.uniformMap[manager.skyShader.Id()]["projection"], 1, false, &manager.matrixes.projection[0])
	opengl.UniformMatrix4fv(manager.uniformMap[manager.skyShader.Id()]["view"], 1, false, &manager.matrixes.view[0])

	manager.lightmappedGenericShader.UseProgram()
	manager.setShader(manager.lightmappedGenericShader.Id())

	//matrixes
	opengl.UniformMatrix4fv(manager.uniformMap[manager.lightmappedGenericShader.Id()]["projection"], 1, false, &manager.matrixes.projection[0])
	opengl.UniformMatrix4fv(manager.uniformMap[manager.lightmappedGenericShader.Id()]["view"], 1, false, &manager.matrixes.view[0])

	opengl.Clear(opengl.COLOR_BUFFER_BIT | opengl.DEPTH_BUFFER_BIT)
}

// Called at the end of a frame
func (manager *Renderer) EndFrame() {
	//if glError := opengl.GetError(); glError != opengl.NO_ERROR {
	//	logger.Error("error: %d\n", glError)
	//}
	//logger.Notice("Calls: %d", numCalls)
	numCalls = 0
}

// Draw the main bsp world
func (manager *Renderer) DrawBsp(world *world.World) {
	modelMatrix := mgl32.Ident4()
	opengl.UniformMatrix4fv(manager.uniformMap[manager.currentShaderId]["model"], 1, false, &modelMatrix[0])
	manager.BindMesh(world.Bsp().Mesh())
	for _, cluster := range world.VisibleClusters() {
		for _, face := range cluster.Faces {
			manager.DrawFace(&face)
		}
	}
	for _, face := range world.Bsp().DefaultCluster().Faces {
		manager.DrawFace(&face)
	}
	for _, cluster := range world.VisibleClusters() {
		for _, prop := range cluster.StaticProps {
			manager.DrawModel(prop.GetModel(), prop.Transform().GetTransformationMatrix())
		}
	}
	for _, prop := range world.Bsp().DefaultCluster().StaticProps {
		manager.DrawModel(prop.GetModel(), prop.Transform().GetTransformationMatrix())
	}
}

// Draw skybox (bsp model, staticprops, sky material)
func (manager *Renderer) DrawSkybox(sky *world.Sky) {
	if sky == nil {
		return
	}

	if sky.GetVisibleBsp() != nil {
		modelMatrix := sky.Transform().GetTransformationMatrix()
		opengl.UniformMatrix4fv(manager.uniformMap[manager.currentShaderId]["model"], 1, false, &modelMatrix[0])
		manager.BindMesh(sky.GetVisibleBsp().Mesh())
		for _, cluster := range sky.GetClusterLeafs() {
			for _, face := range cluster.Faces {
				manager.DrawFace(&face)
			}
		}
		for _, cluster := range sky.GetClusterLeafs() {
			for _, prop := range cluster.StaticProps {
				manager.DrawModel(prop.GetModel(), prop.Transform().GetTransformationMatrix())
			}
		}
	}

	manager.DrawSkyMaterial(sky.GetCubemap())
}

// Render a mesh and its submeshes/primitives
func (manager *Renderer) DrawModel(model *model.Model, transform mgl32.Mat4) {
	opengl.UniformMatrix4fv(manager.uniformMap[manager.currentShaderId]["model"], 1, false, &transform[0])

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
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["baseTextureSampler"], 0)
	}
	// Bind lightmap texture if it exists
	if target.GetLightmap() != nil {
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["useLightmap"], 1)
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["lightmapTextureSampler"], 1)
		target.GetLightmap().Bind()
	} else {
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["useLightmap"], 0)
	}
}

func (manager *Renderer) DrawFace(target *mesh.Face) {
	// Skip materialless faces
	if target.Material() == nil {
		return
	}
	// $basetexture
	target.Material().Bind()
	opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["baseTextureSampler"], 0)

	// Bind lightmap texture if it exists
	if target.IsLightmapped() == true {
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["useLightmap"], 1)
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["lightmapTextureSampler"], 1)
		target.Lightmap().Bind()
	} else {
		opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["useLightmap"], 0)
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

	opengl.CullFace(opengl.FRONT)
	opengl.DepthFunc(opengl.LEQUAL)
	opengl.DepthMask(false)

	manager.skyShader.UseProgram()
	manager.setShader(manager.skyShader.Id())
	opengl.UniformMatrix4fv(manager.uniformMap[manager.skyShader.Id()]["projection"], 1, false, &manager.matrixes.projection[0])
	opengl.UniformMatrix4fv(manager.uniformMap[manager.skyShader.Id()]["view"], 1, false, &manager.matrixes.view[0])

	//DRAW
	skybox.GetMeshes()[0].Bind()
	skybox.GetMeshes()[0].GetMaterial().Bind()
	opengl.Uniform1i(manager.uniformMap[manager.currentShaderId]["cubemapSampler"], 0)
	manager.DrawModel(skybox, mgl32.Ident4())

	// End
	opengl.DepthMask(true)
	opengl.CullFace(uint32(oldCullFaceMode))
	opengl.DepthFunc(uint32(oldDepthFuncMode))

	// Back to default shader
	manager.lightmappedGenericShader.UseProgram()
	manager.setShader(manager.lightmappedGenericShader.Id())
}

// Change the draw format.
func (manager *Renderer) SetWireframeMode(mode bool) {
	if mode == true {
		manager.vertexDrawMode = opengl.LINES
	} else {
		manager.vertexDrawMode = opengl.TRIANGLES
	}
}

func (manager *Renderer) setShader(shader uint32) {
	if manager.currentShaderId != shader {
		manager.currentShaderId = shader
	}
}

func (manager *Renderer) Unregister() {
	manager.skyShader.Destroy()
	manager.lightmappedGenericShader.Destroy()
}

func NewRenderer() *Renderer {
	r := Renderer{}
	r.SetWireframeMode(false)
	r.uniformMap = map[uint32]map[string]int32{}

	return &r
}
