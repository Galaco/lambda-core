package input

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/message/messages"
	"github.com/galaco/Gource-Engine/message/messagetype"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Keyboard key wrapper
type Keyboard struct {
	keysDown [1024]bool
}

// Check if a specific key is pressed
func (keyboard *Keyboard) IsKeyDown(key glfw.Key) bool {
	return keyboard.keysDown[int(key)]
}

// Event manager message receiver.
// Used to catch key events from the window library
func (keyboard *Keyboard) ReceiveMessage(message event.IMessage) {
	switch message.GetType() {
	case messagetype.KeyDown:
		keyboard.keysDown[int(message.(*messages.KeyDown).Key)] = true
	case messagetype.KeyReleased:
		keyboard.keysDown[int(message.(*messages.KeyReleased).Key)] = false
	}
}

func (keyboard *Keyboard) SendMessage() event.IMessage {
	return nil
}

var keyboard Keyboard

func GetKeyboard() *Keyboard {
	return &keyboard
}
