package scene

import (
	"github.com/galaco/Lambda-Core/core/model"
)

type IScene interface {
	Bsp() *model.Bsp
	StaticProps() []model.StaticProp
}
