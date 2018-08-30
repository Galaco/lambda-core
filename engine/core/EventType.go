package core

// Name of an event type in this engines event/messaging
// system
type EventId string

func (id *EventId) String() string {
	return string(*id)
}
