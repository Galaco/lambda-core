package messages

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/core/event"
)

const TypeKeyDown = event.MessageType("KeyDown")

// KeyDown event object for keydown
type KeyDown struct {
	event.Message
	Key keyboard.Key
}

func (message *KeyDown) Type() event.MessageType {
	return TypeKeyDown
}
