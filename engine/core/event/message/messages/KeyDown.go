package messages

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/input/keyboard"
)

// KeyDown event object for keydown
type KeyDown struct {
	event.Message
	Key keyboard.Key
}
