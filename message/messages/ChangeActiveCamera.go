package messages

import (
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/core/interfaces"
)

type ChangeActiveCamera struct {
	event.Message
	Component interfaces.IComponent
}
