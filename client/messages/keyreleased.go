package messages

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/core/event"
)

const TypeKeyReleased = event.MessageType("KeyReleased")

// KeyReleased event object for key released
type KeyReleased struct {
	event.Message
	Key keyboard.Key
}

func (message *KeyReleased) Type() event.MessageType {
	return TypeKeyReleased
}