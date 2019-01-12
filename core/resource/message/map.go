package message

import (
	"github.com/galaco/Lambda-Core/core/event"
	"github.com/galaco/Lambda-Core/core/model"
)

const TypeMapLoaded = event.MessageType("MapLoaded")
const TypeMapUnloaded = event.MessageType("MapUnloaded")

type MapLoaded struct {
	event.Message
	Resource *model.Bsp
}

func (message *MapLoaded) Type() event.MessageType {
	return TypeMapLoaded
}

type MapUnloaded struct {
	event.Message
	Resource *model.Bsp
}

func (message *MapUnloaded) Type() event.MessageType {
	return TypeMapUnloaded
}

func LoadedMap(world *model.Bsp) event.IMessage{
	return &MapLoaded{
		Resource: world,
	}
}

func UnloadedMap(world *model.Bsp) event.IMessage{
	return &MapUnloaded{
		Resource: world,
	}
}