package messages

import (
	"github.com/galaco/go-me-engine/engine/event"
)

type MouseMove struct {
	event.Message
	X float64
	Y float64
}
