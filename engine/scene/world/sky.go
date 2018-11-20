package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/go-gl/mathgl/mgl32"
)

type Sky struct {
	geometry  *model.Bsp
	skybox    *model.Model
	props     []*StaticProp
	transform entity.Transform
}

func (sky *Sky) GetVisibleBsp() *model.Bsp {
	return sky.geometry
}

func (sky *Sky) GetBackdrop() *model.Model {
	return sky.skybox
}

func (sky *Sky) GetVisibleProps() []*StaticProp {
	return sky.props
}

func (sky *Sky) Transform() *entity.Transform {
	return &sky.transform
}

func NewSky(model *model.Bsp, sky *model.Model, props []*StaticProp, position mgl32.Vec3, scale float32) *Sky {
	s := Sky{
		geometry: model,
		skybox:   sky,
	}

	skyCameraPosition := (mgl32.Vec3{0, 0, 0}).Sub(position)
	skyCameraScale := mgl32.Vec3{scale, scale, scale}

	s.transform.Position = skyCameraPosition.Mul(scale)
	s.transform.Scale = skyCameraScale

	// remap prop transform to real world
	for _, prop := range props {
		prop.Transform().Position = prop.Transform().Position.Add(skyCameraPosition)
		prop.Transform().Position = prop.Transform().Position.Mul(scale)
		prop.Transform().Scale = skyCameraScale
	}

	s.props = props
	return &s
}
