package event

// Types that want to be able to receive events from the
// event manager should implement this
type IEventListenable interface {
	ReceiveMessage(IMessage)
}
