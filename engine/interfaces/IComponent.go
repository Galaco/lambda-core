package interfaces

import "github.com/galaco/go-me-engine/engine/core"

// Component interface
// All components need to implement this
type IComponent interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	Initialize()
	GetOwnerHandle() core.Handle
	SetOwnerHandle(core.Handle)
	Update(float64)
	Destroy()
}
