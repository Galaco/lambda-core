package event

import "github.com/galaco/go-me-engine/engine/core"

type Message struct {
	Type core.EventId
}

func (message *Message) SetType(messageType core.EventId) {
	message.Type = messageType
}

func (message Message) GetType() core.EventId {
	return message.Type
}