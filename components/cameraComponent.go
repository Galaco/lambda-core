package components

import (
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/engine/input"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

const cameraSpeed = float64(100)
const sensitivity = float64(0.05)

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
	vel := cameraSpeed // * dt
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

	component.owner.GetTransformComponent().Rotation[0] += float32(input.GetMouse().GetCoordinates()[0] * sensitivity) // * dt)
	component.owner.GetTransformComponent().Rotation[1] += float32(input.GetMouse().GetCoordinates()[1] * sensitivity) // * dt)

	component.updateVectors()
}

// Update the camera directional properties with any changes
func (component *CameraComponent) updateVectors() {
	rot := component.owner.GetTransformComponent().Rotation
	//rot[0] = YAW, rot[1] = PITCH

	// Calculate the new Front vector
	component.Direction = mgl32.Vec3{
		float32(math.Cos(float64(rot[1])) * math.Sin(float64(rot[0]))),
		float32(math.Sin(float64(rot[1]))),
		float32(math.Cos(float64(rot[1])) * math.Cos(float64(rot[0]))),
	}
	// Also re-calculate the Right and Up vector
	component.Right = mgl32.Vec3{
		float32(math.Sin(float64(rot[0]) - math.Pi/2)),
		0,
		float32(math.Cos(float64(rot[0]) - math.Pi/2)),
	}
	component.Up = component.Right.Cross(component.Direction)
}

func NewCameraComponent() *CameraComponent {
	c := CameraComponent{}
	c.Etype = T_CameraComponent

	return &c
}
