package input

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/go-gl/mathgl/mgl64"
)

type Mouse struct {
	previousPosition mgl64.Vec2
	position mgl64.Vec2
}

func (mouse *Mouse) GetCoordinates() mgl64.Vec2 {
	return mouse.position
}

func (mouse *Mouse) GetPreviousCoordinates() mgl64.Vec2 {
	return mouse.previousPosition
}

func (mouse *Mouse) GetChange() mgl64.Vec2 {
	return mouse.previousPosition.Sub(mouse.position)
}

func (mouse *Mouse) ReceiveMessage(message interfaces.IMessage) {
	if message.GetType() == messagetype.MouseMove {
		msg := message.(*messages.MouseMove)
		mouse.previousPosition = mouse.position
		mouse.position[0] = msg.X
		mouse.position[1] = msg.Y
	}
}

func (mouse *Mouse) Update()  {
	mouse.previousPosition = mouse.position
}

func (mouse *Mouse) SendMessage() interfaces.IMessage {
	return nil
}



var mouse Mouse

func GetMouse() *Mouse {
	return &mouse
}