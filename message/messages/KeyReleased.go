package messages

import (
	"github.com/galaco/Gource-Engine/engine/event"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type KeyReleased struct {
	event.Message
	Key glfw.Key
}
