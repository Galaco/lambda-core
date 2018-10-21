package messages

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type KeyDown struct {
	event.Message
	Key glfw.Key
}
