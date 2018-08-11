package interfaces

type IMessage interface {
	SetType(string)
	GetType() string
}