package messages

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/core/event"
)

// KeyReleased event object for key released
type KeyReleased struct {
	event.Message
	Key keyboard.Key
}
