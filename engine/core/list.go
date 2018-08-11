package core

type EventId string

func (id *EventId) String() string {
	return string(*id)
}