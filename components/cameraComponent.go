package components

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/message/input"
	"log"
	"github.com/galaco/go-me-engine/engine/entity"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type CameraComponent struct {
	entity.Component
}

func (component *CameraComponent) Initialize() {
	event.GetEventManager().RegisterEvent(event.KeyDown, component)
	event.GetEventManager().RegisterEvent(event.KeyHeld, component)
	event.GetEventManager().RegisterEvent(event.KeyReleased, component)
}

func (component *CameraComponent) ReceiveMessage(message interfaces.IMessage) {
	switch message.GetType() {
	case "KeyDown":
		m := message.(*input.KeyDown)
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
	case "KeyHeld":
		m := message.(*input.KeyHeld)
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
	case "KeyReleased":
		m := message.(*input.KeyReleased)
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