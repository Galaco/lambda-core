package event

import "github.com/galaco/Gource-Engine/engine/core/event/message"

// IEventListenable Types that want to be able to receive events from the
// event manager should implement this
type IEventListenable interface {
	ReceiveMessage(message message.IMessage)
}
