package components

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/galaco/go-me-engine/engine/factory"
	"math"
	"github.com/galaco/go-me-engine/engine/input"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/message/messages"
)

const cameraSpeed = float64(25)
const sensitivity = float64(0.1)

type CameraComponent struct {
	base.Component
	Up mgl32.Vec3
	Right mgl32.Vec3
	Front mgl32.Vec3
	worldUp mgl32.Vec3
	owner *base.Entity
	frameTime float64
}

func (component *CameraComponent) Initialize() {
	component.owner = factory.GetObjectManager().GetEntityByHandle(component.GetOwnerHandle()).(*base.Entity)
	component.Up = mgl32.Vec3{0, 1, 0}
	component.worldUp = mgl32.Vec3{0, 1, 0}
	component.Front = mgl32.Vec3{0, 0, -1}

	event.GetEventManager().Dispatch(messagetype.ChangeActiveCamera, &messages.ChangeActiveCamera{Component: component})
}

func (component *CameraComponent) ReceiveMessage(message interfaces.IMessage) {

}

func (component *CameraComponent) Update(dt float64) {
	vel := cameraSpeed * dt
	if input.GetKeyboard().IsKeyDown(glfw.KeyA) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Sub(component.Right.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(glfw.KeyW) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Add(component.Front.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(glfw.KeyS) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Sub(component.Front.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(glfw.KeyD) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Add(component.Right.Mul(float32(vel)))
	}
	//component.owner.GetTransformComponent().Rotation[0] += float32(input.GetMouse().GetChange()[0] * sensitivity)
	//component.owner.GetTransformComponent().Rotation[1] += float32(input.GetMouse().GetChange()[1] * sensitivity)

	component.updateVectors()
}

func (component *CameraComponent) updateVectors() {
	// Calculate the new Front vector
	front := mgl32.Vec3{}
	rot := component.owner.GetTransformComponent().Rotation
	front[0] = float32(math.Cos(float64(mgl32.DegToRad(rot[1]))) * math.Cos(float64(mgl32.DegToRad(rot[0]))))
	front[1] = float32(math.Sin(float64(mgl32.DegToRad(rot[0]))))
	front[2] = float32(math.Sin(float64(mgl32.DegToRad(rot[1]))) * math.Cos(float64(mgl32.DegToRad(rot[0]))))
	component.Front = front.Normalize()
	// Also re-calculate the Right and Up vector
	component.Right = component.Front.Cross(component.worldUp).Normalize() // Normalize the vectors, because their length gets closer to 0 the more you look up or down which results in slower movement.
	component.Up    = component.Right.Cross(component.Front).Normalize()
}

func NewCameraComponent() *CameraComponent{
	c := &CameraComponent{}
	c.Etype = T_CameraComponent

	return c
}