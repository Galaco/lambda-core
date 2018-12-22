package messages

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/core/event"
)

// KeyHeld event object for when key is held down
type KeyHeld struct {
	event.Message
	Key keyboard.Key
}
