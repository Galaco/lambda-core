package messages

import (
	"github.com/galaco/Gource-Engine/engine/event"
	"github.com/galaco/Gource-Engine/engine/interfaces"
)

type ChangeActiveCamera struct {
	event.Message
	Component interfaces.IComponent
}
