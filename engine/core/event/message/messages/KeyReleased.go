package messages

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/input/keyboard"
)

// KeyReleased event object for key released
type KeyReleased struct {
	event.Message
	Key keyboard.Key
}
