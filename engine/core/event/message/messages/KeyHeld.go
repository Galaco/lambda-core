package messages

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/input/keyboard"
)

// KeyHeld event object for when key is held down
type KeyHeld struct {
	event.Message
	Key keyboard.Key
}
