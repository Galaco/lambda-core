package interfaces

import (
	"github.com/galaco/bsp-viewer/engine/core"
)

type IEntity interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	GetComponents() []core.Handle
	AddComponent(handle core.Handle)
}
