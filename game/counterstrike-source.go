package game

import (
	"github.com/galaco/Gource-Engine/engine/loader/entity/classmap"
	"github.com/galaco/Gource-Engine/entity/common"
)

type CounterstrikeSource struct {

}

func (target *CounterstrikeSource) RegisterEntityClasses() {
	loader.RegisterClass(&common.PropDynamic{})
	loader.RegisterClass(&common.PropDynamicOrnament{})
	loader.RegisterClass(&common.PropDynamicOverride{})
	loader.RegisterClass(&common.PropPhysics{})
	loader.RegisterClass(&common.PropPhysicsMultiplayer{})
	loader.RegisterClass(&common.PropPhysicsOverride{})
	loader.RegisterClass(&common.PropRagdoll{})
}