package event

import "github.com/galaco/Gource-Engine/engine/core/event/message"

// Event Queue item.
// Contains the event name, and a message
type QueueItem struct {
	EventName message.Id
	Message   message.IMessage
}
