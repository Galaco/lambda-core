package messages

import (
	"github.com/galaco/Gource/engine/event"
	"github.com/galaco/Gource/engine/interfaces"
)

type ChangeActiveCamera struct {
	event.Message
	Component interfaces.IComponent
}
