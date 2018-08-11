package event

import "github.com/galaco/go-me-engine/engine/interfaces"

// Event Queue item.
// Contains the event name,
type QueueItem struct {
	EventName Id
	Message interfaces.IMessage
}
