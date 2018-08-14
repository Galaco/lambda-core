package input

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/go-gl/mathgl/mgl64"
)

type Mouse struct {
	previousPosition mgl64.Vec2
	change mgl64.Vec2
}

func (mouse *Mouse) GetCoordinates() mgl64.Vec2 {
	return mouse.change
}

func (mouse *Mouse) ReceiveMessage(message interfaces.IMessage) {
	if message.GetType() == messagetype.MouseMove {
		msg := message.(*messages.MouseMove)
		mouse.change[0] = msg.X
		mouse.change[1] = msg.Y
	}
}

func (mouse *Mouse) Update()  {
	mouse.change[0] = 0
	mouse.change[1] = 0
}

func (mouse *Mouse) SendMessage() interfaces.IMessage {
	return nil
}



var mouse Mouse

func GetMouse() *Mouse {
	return &mouse
}