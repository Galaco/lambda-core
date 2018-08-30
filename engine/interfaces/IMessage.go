package interfaces

import "github.com/galaco/go-me-engine/engine/core"

// Generic event manager message interface
// All messages need to implement this
type IMessage interface {
	SetType(core.EventId)
	GetType() core.EventId
}
