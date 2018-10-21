package messages

import (
	"github.com/galaco/Gource-Engine/engine/component"
	"github.com/galaco/Gource-Engine/engine/core/event"
)

type ChangeActiveCamera struct {
	event.Message
	Component component.IComponent
}
