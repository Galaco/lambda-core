package event

import "github.com/galaco/Gource-Engine/engine/core/event/message"

// Message is a generic event message
// Contains the type of event
type Message struct {
	Type message.Id
}

// SetType Sets the type of this event
func (message *Message) SetType(messageType message.Id) {
	message.Type = messageType
}

// GetType Gets type of event
func (message Message) GetType() message.Id {
	return message.Type
}
