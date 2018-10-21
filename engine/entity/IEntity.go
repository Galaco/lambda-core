package entity

import (
	"github.com/galaco/Gource-Engine/engine/core"
)

// Entity interface
// All game entities need to implement this
type IEntity interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	GetComponents() []core.Handle
	AddComponent(handle core.Handle)
}
