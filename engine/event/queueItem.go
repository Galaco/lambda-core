package event

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/Gource-Engine/engine/interfaces"
)

// Event Queue item.
// Contains the event name, and a message
type QueueItem struct {
	EventName core.EventId
	Message   interfaces.IMessage
}
