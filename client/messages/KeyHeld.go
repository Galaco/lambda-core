package messages

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/core/event"
)

const TypeKeyHeld = event.MessageType("KeyHeld")

// KeyHeld event object for when key is held down
type KeyHeld struct {
	event.Message
	Key keyboard.Key
}

func (message *KeyHeld) Type() event.MessageType {
	return TypeKeyHeld
}
