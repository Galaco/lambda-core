package messages

import (
	"github.com/galaco/Gource-Engine/engine/event"
)

type MouseMove struct {
	event.Message
	X float64
	Y float64
}
