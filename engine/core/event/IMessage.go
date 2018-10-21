package event

// Generic event manager message interface
// All messages need to implement this
type IMessage interface {
	SetType(Id)
	GetType() Id
}
