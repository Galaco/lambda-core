package messages

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
)

type MouseMove struct {
	event.Message
	X float64
	Y float64
}
