package interfaces

import (
	"github.com/galaco/go-me-engine/engine/core"
)

type IEntity interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	GetComponents() []core.Handle
	AddComponent(handle core.Handle)
}
