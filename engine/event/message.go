package event

type Message struct {
	Type string
}

func (message *Message) SetType(messageType string) {
	message.Type = messageType
}

func (message Message) GetType() string {
	return message.Type
}