package messages

import (
	"github.com/galaco/Gource/engine/event"
)

type MouseMove struct {
	event.Message
	X float64
	Y float64
}
