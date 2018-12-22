package messages

import (
	"github.com/galaco/Gource-Engine/core/event"
)

// MouseMove event object for when mouse is moved
type MouseMove struct {
	event.Message
	X float64
	Y float64
}
