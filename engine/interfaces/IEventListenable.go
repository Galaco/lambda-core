package interfaces


type IEventListenable interface {
	ReceiveMessage(IMessage)
}