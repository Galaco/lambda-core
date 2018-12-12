package message

// Id Name of an event type in this engines event/messaging
// system
type Id string

// String returns Id as a string
func (id *Id) String() string {
	return string(*id)
}
