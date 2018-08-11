package base

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/core"
)

type Component struct {
	handle core.Handle
	owner core.Handle
	Etype core.EType
}

func (component *Component) SetHandle(handle core.Handle) {
	component.handle = handle
}

func (component *Component) Initialize() {

}

func (component *Component) GetType() core.EType {
	return component.Etype
}

func (component *Component) GetHandle() core.Handle{
	return component.handle
}

func (component *Component) GetOwnerHandle() core.Handle{
	return component.owner
}

func (component *Component) SetOwnerHandle(handle core.Handle) {
	component.owner = handle
}

func (component *Component) ReceiveMessage(message interfaces.IMessage) {
}

func (component *Component) SendMessage() interfaces.IMessage {
	return nil
}

func (component *Component) Update(dt float64) {
}

func (component *Component) Destroy() {
}
