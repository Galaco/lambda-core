package components

import (
	"github.com/galaco/Gource-Engine/engine/base"
	"github.com/galaco/Gource-Engine/engine/event"
	"github.com/galaco/Gource-Engine/engine/factory"
	"github.com/galaco/Gource-Engine/engine/input"
	"github.com/galaco/Gource-Engine/engine/interfaces"
	"github.com/galaco/Gource-Engine/message/messages"
	"github.com/galaco/Gource-Engine/message/messagetype"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

const cameraSpeed = float64(320)
const sensitivity = float64(0.03)

var minVerticalRotation = mgl32.DegToRad(90)
var maxVerticalRotation = mgl32.DegToRad(270)

type CameraComponent struct {
	base.Component
	Up        mgl32.Vec3
	Right     mgl32.Vec3
	Direction mgl32.Vec3
	worldUp   mgl32.Vec3
	owner     *base.Entity
	frameTime float64
}

func (component *CameraComponent) Initialize() {
	component.owner = factory.GetObjectManager().GetEntityByHandle(component.GetOwnerHandle()).(*base.Entity)
	component.Up = mgl32.Vec3{0, 1, 0}
	component.worldUp = mgl32.Vec3{0, 1, 0}
	component.Direction = mgl32.Vec3{0, 0, -1}

	event.GetEventManager().Dispatch(messagetype.ChangeActiveCamera, &messages.ChangeActiveCamera{Component: component})
}

func (component *CameraComponent) ReceiveMessage(message interfaces.IMessage) {

}

func (component *CameraComponent) Update(dt float64) {
	vel := cameraSpeed * dt
	if input.GetKeyboard().IsKeyDown(glfw.KeyW) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Add(component.Direction.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(glfw.KeyA) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Sub(component.Right.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(glfw.KeyS) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Sub(component.Direction.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(glfw.KeyD) {
		component.owner.GetTransformComponent().Position = component.owner.GetTransformComponent().Position.Add(component.Right.Mul(float32(vel)))
	}

	component.owner.GetTransformComponent().Rotation[0] -= float32(input.GetMouse().GetCoordinates()[0] * sensitivity)
	component.owner.GetTransformComponent().Rotation[2] -= float32(input.GetMouse().GetCoordinates()[1] * sensitivity)

	// Lock vertical rotation
	if component.owner.GetTransformComponent().Rotation[2] > maxVerticalRotation {
		component.owner.GetTransformComponent().Rotation[2] = maxVerticalRotation
	}
	if component.owner.GetTransformComponent().Rotation[2] < minVerticalRotation {
		component.owner.GetTransformComponent().Rotation[2] = minVerticalRotation
	}

	component.updateVectors()
}

// Update the camera directional properties with any changes
func (component *CameraComponent) updateVectors() {
	rot := component.owner.GetTransformComponent().Rotation

	// Calculate the new Front vector
	component.Direction = mgl32.Vec3{
		float32(math.Cos(float64(rot[2])) * math.Sin(float64(rot[0]))),
		float32(math.Cos(float64(rot[2])) * math.Cos(float64(rot[0]))),
		float32(math.Sin(float64(rot[2]))),
	}
	// Also re-calculate the Right and Up vector
	component.Right = mgl32.Vec3{
		float32(math.Sin(float64(rot[0]) - math.Pi/2)),
		float32(math.Cos(float64(rot[0]) - math.Pi/2)),
		0,
	}
	component.Up = component.Right.Cross(component.Direction)
}

func NewCameraComponent() *CameraComponent {
	c := CameraComponent{}

	return &c
}
