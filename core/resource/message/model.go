package message

import (
	"github.com/galaco/Lambda-Core/core/event"
	"github.com/galaco/Lambda-Core/core/model"
)

const TypeModelLoaded = event.MessageType("ModelLoaded")
const TypeModelUnloaded = event.MessageType("ModelUnloaded")

type PropLoaded struct {
	event.Message
	Resource *model.Model
}

func (message *PropLoaded) Type() event.MessageType {
	return TypeModelLoaded
}

type PropUnloaded struct {
	event.Message
	Resource *model.Model
}

func (message *PropUnloaded) Type() event.MessageType {
	return TypeModelUnloaded
}

func LoadedModel(mod *model.Model) event.IMessage{
	return &PropLoaded{
		Resource: mod,
	}
}

func UnloadedModel(mod *model.Model) event.IMessage{
	return &PropUnloaded{
		Resource: mod,
	}
}