package input

import (
	"github.com/galaco/Gource-Engine/client/messages"
	"github.com/galaco/Gource-Engine/core/event"
	"github.com/go-gl/mathgl/mgl32"
)

// Mouse information, about change from previous poll.
// Note: Mouse is a struct containing mouse information, it doesn't have
// any direct interaction with the window
type Mouse struct {
	change mgl32.Vec2
}

// GetCoordinates return current mouse position
func (mouse *Mouse) GetCoordinates() mgl32.Vec2 {
	return mouse.change
}

// ReceiveMessage mouse receives updated info from the event queue about
// mouse interaction
func (mouse *Mouse) ReceiveMessage(message event.IMessage) {
	if message.GetType() == messages.TypeMouseMove {
		msg := message.(*messages.MouseMove)
		mouse.change[0] = float32(msg.X)
		mouse.change[1] = float32(msg.Y)
	}
}

// Update The Mouse should be reset to screen center
func (mouse *Mouse) Update() {
	mouse.change[0] = 0
	mouse.change[1] = 0
}

func (mouse *Mouse) SendMessage() event.IMessage {
	return nil
}

var mouse Mouse

// GetMouse return static mouse
func GetMouse() *Mouse {
	return &mouse
}
