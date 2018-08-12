package messages

import (
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/interfaces"
)

type ChangeActiveCamera struct {
	event.Message
	Component interfaces.IComponent
}
