package input

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/message/messages"
	"github.com/galaco/Gource-Engine/message/messagetype"
	"github.com/go-gl/mathgl/mgl64"
)

// Mouse information, about change from previous poll
type Mouse struct {
	change mgl64.Vec2
}

func (mouse *Mouse) GetCoordinates() mgl64.Vec2 {
	return mouse.change
}

func (mouse *Mouse) ReceiveMessage(message event.IMessage) {
	if message.GetType() == messagetype.MouseMove {
		msg := message.(*messages.MouseMove)
		mouse.change[0] = msg.X
		mouse.change[1] = msg.Y
	}
}

func (mouse *Mouse) Update() {
	mouse.change[0] = 0
	mouse.change[1] = 0
}

func (mouse *Mouse) SendMessage() event.IMessage {
	return nil
}

var mouse Mouse

func GetMouse() *Mouse {
	return &mouse
}
