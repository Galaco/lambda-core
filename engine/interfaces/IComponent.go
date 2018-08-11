package interfaces

import "github.com/galaco/bsp-viewer/engine/core"

type IComponent interface {
	SetHandle(core.Handle)
	GetHandle() core.Handle
	Initialize()
	GetOwnerHandle() core.Handle
	SetOwnerHandle(core.Handle)
	ReceiveMessage(IMessage)
	SendMessage() IMessage
	Update(float64)
	Destroy()
}
