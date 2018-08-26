package interfaces

import (
	"github.com/galaco/go-me-engine/engine/core"
)

// Entity interface
// All game entities need to implement this
type IEntity interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	GetComponents() []core.Handle
	AddComponent(handle core.Handle)
}
