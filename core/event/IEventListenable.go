package event

// IEventListenable Types that want to be able to receive events from the
// event manager should implement this
type IEventListenable interface {
	ReceiveMessage(message IMessage)
}
