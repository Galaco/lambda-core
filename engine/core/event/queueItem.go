package event

// Event Queue item.
// Contains the event name, and a message
type QueueItem struct {
	EventName Id
	Message   IMessage
}
