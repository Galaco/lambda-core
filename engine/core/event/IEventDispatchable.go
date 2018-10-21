package event

// Types that can dispatch event to the event manager
// should implement this
type IEventDispatchable interface {
	SendMessage() IMessage
}
