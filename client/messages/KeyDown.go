package messages

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/core/event"
)

// KeyDown event object for keydown
type KeyDown struct {
	event.Message
	Key keyboard.Key
}
