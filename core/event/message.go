package event

// Message is a generic event message
// Contains the type of event
type Message struct {
	Type MessageType
}

// SetType Sets the type of this event
func (message *Message) SetType(messageType MessageType) {
	message.Type = messageType
}

// GetType Gets type of event
func (message Message) GetType() MessageType {
	return message.Type
}
