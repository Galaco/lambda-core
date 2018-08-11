package interfaces

import "github.com/galaco/go-me-engine/engine/core"

type IMessage interface {
	SetType(core.EventId)
	GetType() core.EventId
}