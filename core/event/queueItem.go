package event

// QueueItem Event Queue item.
// Contains the event name, and a message
type QueueItem struct {
	EventName MessageType
	Message   IMessage
}
