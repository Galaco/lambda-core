package event

import (
	"github.com/galaco/go-me-engine/engine/core"
	"github.com/galaco/go-me-engine/engine/interfaces"
)

// Event Queue item.
// Contains the event name, and a message
type QueueItem struct {
	EventName core.EventId
	Message   interfaces.IMessage
}
