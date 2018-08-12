package interfaces


type IEventDispatchable interface {
	SendMessage() IMessage
}