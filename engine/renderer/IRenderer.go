package renderer

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/scene/world"
	"github.com/go-gl/mathgl/mgl32"
)

type IRenderer interface {
	StartFrame(*entity.Camera)
	LoadShaders()
	DrawBsp(*world.VisibleWorld)
	DrawSkybox(*world.Sky)
	DrawStaticProps([]*world.StaticProp)
	DrawModel(*model.Model, mgl32.Mat4)
	DrawSkyMaterial(*model.Model)
	SetWireframeMode(bool)
	EndFrame()
}
