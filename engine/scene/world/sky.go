package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/go-gl/mathgl/mgl32"
)

type Sky struct {
	skybox *model.Model
	props []*StaticProp
	transform entity.Transform
}

func (sky *Sky) GetVisibleBsp() *model.Model {
	return sky.skybox
}

func (sky *Sky) GetVisibleProps() []*StaticProp {
	return sky.props
}

func (sky *Sky) Transform() *entity.Transform {
	return &sky.transform
}

func NewSky(model *model.Model, props []*StaticProp, position mgl32.Vec3, scale float32) *Sky {
	s := Sky{
		skybox: model,
	}

	skyCameraPosition := position.Mul(-1)
	skyCameraScale := mgl32.Vec3{scale, scale, scale}

	s.transform.Position = skyCameraPosition
	s.transform.Scale = skyCameraScale

	// remap prop transform to real world
	for _,prop := range props {
		prop.Transform().Position = prop.Transform().Position.Add(skyCameraPosition)
		prop.Transform().Scale = skyCameraScale
	}

	s.props = props
	return &s
}
