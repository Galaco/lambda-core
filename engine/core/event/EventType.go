package event

// Name of an event type in this engines event/messaging
// system
type Id string

func (id *Id) String() string {
	return string(*id)
}
