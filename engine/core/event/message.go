package event

// A generic event message
// Contains the type of event
type Message struct {
	Type Id
}

// Set the type of this event
func (message *Message) SetType(messageType Id) {
	message.Type = messageType
}

// Get type of event
func (message Message) GetType() Id {
	return message.Type
}
