package message

import (
	"github.com/galaco/Lambda-Core/core/event"
	"github.com/galaco/Lambda-Core/core/material"
)

const TypeMaterialLoaded = event.MessageType("MaterialLoaded")
const TypeMaterialUnloaded = event.MessageType("MaterialUnloaded")

type MaterialLoaded struct {
	event.Message
	Resource material.IMaterial
}

func (message *MaterialLoaded) Type() event.MessageType {
	return TypeMaterialLoaded
}

type MaterialUnloaded struct {
	event.Message
	Resource material.IMaterial
}

func (message *MaterialUnloaded) Type() event.MessageType {
	return TypeMaterialUnloaded
}

func LoadedMaterial(mat material.IMaterial) event.IMessage{
	return &MaterialLoaded{
		Resource: mat,
	}
}

func UnloadedMaterial(mat material.IMaterial) event.IMessage{
	return &MaterialUnloaded{
		Resource: mat,
	}
}