package scene

import (
	"github.com/galaco/Lambda-Core/core/model"
)

// IScene
type IScene interface {
	// Bsp
	Bsp() *model.Bsp
	// StaticProps
	StaticProps() []model.StaticProp
}
