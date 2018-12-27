package scene

import (
	"github.com/galaco/Gource-Engine/core/model"
)

type Scene struct {
	bsp model.Bsp
	staticProps []model.StaticProp
}


func (s *Scene) Bsp() *model.Bsp {
	return &s.bsp
}

func (s *Scene) StaticProps() []model.StaticProp {
	return s.staticProps
}

func NewScene(bsp model.Bsp, staticProps []model.StaticProp) *Scene {
	return &Scene{
		bsp: bsp,
		staticProps: staticProps,
	}
}