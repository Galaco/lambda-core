package input

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/message/messages"
)

type Keyboard struct {
	keysDown [1024]bool
}

func (keyboard *Keyboard) IsKeyDown(key glfw.Key) bool {
	return keyboard.keysDown[int(key)]
}

func (keyboard *Keyboard) ReceiveMessage(message interfaces.IMessage) {
	switch message.GetType() {
	case messagetype.KeyDown:
		keyboard.keysDown[int(message.(*messages.KeyDown).Key)] = true
	case messagetype.KeyReleased:
		keyboard.keysDown[int(message.(*messages.KeyReleased).Key)] = false
	}
}

func (keyboard *Keyboard) SendMessage() interfaces.IMessage {
	return nil
}

var keyboard Keyboard

func GetKeyboard() *Keyboard {
	return &keyboard
}