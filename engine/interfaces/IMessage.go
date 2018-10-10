package interfaces

import "github.com/galaco/Gource-Engine/engine/core"

// Generic event manager message interface
// All messages need to implement this
type IMessage interface {
	SetType(core.EventId)
	GetType() core.EventId
}
