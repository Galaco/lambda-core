package message

import (
	"github.com/galaco/Lambda-Core/core/event"
	"github.com/galaco/Lambda-Core/core/texture"
)

const TypeTextureLoaded = event.MessageType("TextureLoaded")
const TypeTextureUnloaded = event.MessageType("TextureUnloaded")

type TextureLoaded struct {
	event.Message
	Resource texture.ITexture
}

func (message *TextureLoaded) Type() event.MessageType {
	return TypeTextureLoaded
}

type TextureUnloaded struct {
	event.Message
	Resource texture.ITexture
}

func (message *TextureUnloaded) Type() event.MessageType {
	return TypeTextureUnloaded
}

func LoadedTexture(tex texture.ITexture) event.IMessage {
	return &TextureLoaded{
		Resource: tex,
	}
}

func UnloadedTexture(tex texture.ITexture) event.IMessage {
	return &TextureUnloaded{
		Resource: tex,
	}
}
