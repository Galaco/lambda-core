package event

import "github.com/galaco/Gource-Engine/engine/core"

// A generic event message
// Contains the type of event
type Message struct {
	Type core.EventId
}

// Set the type of this event
func (message *Message) SetType(messageType core.EventId) {
	message.Type = messageType
}

// Get type of event
func (message Message) GetType() core.EventId {
	return message.Type
}
