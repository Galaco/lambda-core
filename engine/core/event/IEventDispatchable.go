package event

import "github.com/galaco/Gource-Engine/engine/core/event/message"

// Types that can dispatch event to the event manager
// should implement this
type IEventDispatchable interface {
	SendMessage() message.IMessage
}
