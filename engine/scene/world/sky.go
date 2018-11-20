package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/go-gl/mathgl/mgl32"
)

type Sky struct {
	geometry  *model.Bsp
	clusterLeafs    []*model.ClusterLeaf
	transform entity.Transform
}

func (sky *Sky) GetVisibleBsp() *model.Bsp {
	return sky.geometry
}

func (sky *Sky) GetClusterLeafs() []*model.ClusterLeaf {
	return sky.clusterLeafs
}

func (sky *Sky) Transform() *entity.Transform {
	return &sky.transform
}

func NewSky(model *model.Bsp, clusterLeafs []*model.ClusterLeaf, position mgl32.Vec3, scale float32) *Sky {
	s := Sky{
		geometry: model,
		clusterLeafs: clusterLeafs,
	}

	skyCameraPosition := (mgl32.Vec3{0, 0, 0}).Sub(position)
	skyCameraScale := mgl32.Vec3{scale, scale, scale}

	s.transform.Position = skyCameraPosition.Mul(scale)
	s.transform.Scale = skyCameraScale

	// remap prop transform to real world
	for _,l := range s.clusterLeafs{
		for _, prop := range l.StaticProps {
			prop.Transform().Position = prop.Transform().Position.Add(skyCameraPosition)
			prop.Transform().Position = prop.Transform().Position.Mul(scale)
			prop.Transform().Scale = skyCameraScale
		}
	}
	return &s
}
