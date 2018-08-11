package components

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/message/messages"
	"log"
	"github.com/galaco/go-me-engine/engine/entity"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/galaco/go-me-engine/message/messagetype"
)

type CameraComponent struct {
	entity.Component
}

func (component *CameraComponent) Initialize() {
	event.GetEventManager().Listen(messagetype.KeyDown, component)
	event.GetEventManager().Listen(messagetype.KeyHeld, component)
	event.GetEventManager().Listen(messagetype.KeyReleased, component)
}

func (component *CameraComponent) ReceiveMessage(message interfaces.IMessage) {
	switch message.GetType() {
	case messagetype.KeyDown:
		m := message.(*messages.KeyDown)
		switch m.Key {
		case glfw.KeyA:
			log.Println("A pressed")
		case glfw.KeyW:
			log.Println("W pressed")
		case glfw.KeyS:
			log.Println("S pressed")
		case glfw.KeyD:
			log.Println("D pressed")
		}
	case messagetype.KeyHeld:
		m := message.(*messages.KeyHeld)
		switch m.Key {
		case glfw.KeyA:
			log.Println("A held")
		case glfw.KeyW:
			log.Println("W held")
		case glfw.KeyS:
			log.Println("S held")
		case glfw.KeyD:
			log.Println("D held")
		}
	case messagetype.KeyReleased:
		m := message.(*messages.KeyReleased)
		switch m.Key {
		case glfw.KeyA:
			log.Println("A released")
		case glfw.KeyW:
			log.Println("W released")
		case glfw.KeyS:
			log.Println("S released")
		case glfw.KeyD:
			log.Println("D released")
		}
	}
}

func (component *CameraComponent) Update(dt float64) {

}