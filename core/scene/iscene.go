package scene

import (
	"github.com/galaco/Gource-Engine/core/model"
)

type IScene interface {
	Bsp() *model.Bsp
	StaticProps() []model.StaticProp
}
