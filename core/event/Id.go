package event

type MessageType string

type EventHandle uint

var eventHandleCounter EventHandle

func newEventHandle() EventHandle {
	eventHandleCounter++
	return eventHandleCounter
}
