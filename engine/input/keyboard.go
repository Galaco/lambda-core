package input

import (
	"github.com/galaco/Gource-Engine/engine/core/event/message"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messages"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messagetype"
	"github.com/galaco/Gource-Engine/engine/input/keyboard"
)

// Keyboard key wrapper
type Keyboard struct {
	keysDown [1024]bool
}

// IsKeyDown Check if a specific key is pressed
func (keyboard *Keyboard) IsKeyDown(key keyboard.Key) bool {
	return keyboard.keysDown[int(key)]
}

// ReceiveMessage Event manager message receiver.
// Used to catch key events from the window library
func (keyboard *Keyboard) ReceiveMessage(message message.IMessage) {
	switch message.GetType() {
	case messagetype.KeyDown:
		keyboard.keysDown[int(message.(*messages.KeyDown).Key)] = true
	case messagetype.KeyReleased:
		keyboard.keysDown[int(message.(*messages.KeyReleased).Key)] = false
	}
}

func (keyboard *Keyboard) SendMessage() message.IMessage {
	return nil
}

var staticKeyboard Keyboard

// GetKeyboard return static keyboard
func GetKeyboard() *Keyboard {
	return &staticKeyboard
}
