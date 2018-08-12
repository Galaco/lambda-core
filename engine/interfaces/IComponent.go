package interfaces

import "github.com/galaco/go-me-engine/engine/core"

type IComponent interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	GetType() core.EType
	Initialize()
	GetOwnerHandle() core.Handle
	SetOwnerHandle(core.Handle)
	Update(float64)
	Destroy()
}
