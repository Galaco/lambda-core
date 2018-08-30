package camera

import (
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	currentCameraComponent *components.CameraComponent
	owner                  *base.Entity
}

func (camera *Camera) Initialize() {
	event.GetEventManager().Listen(messagetype.ChangeActiveCamera, camera)
}

func (camera *Camera) ReceiveMessage(message interfaces.IMessage) {
	if message.GetType() == messagetype.ChangeActiveCamera {
		camera.currentCameraComponent = message.(*messages.ChangeActiveCamera).Component.(*components.CameraComponent)
		camera.owner = factory.GetObjectManager().GetEntityByHandle(camera.currentCameraComponent.GetOwnerHandle()).(*base.Entity)
	}
}

func (camera *Camera) SendMessage() interfaces.IMessage {
	return nil
}

func (camera *Camera) Update(dt float64) {
	camera.currentCameraComponent.Update(dt)
}

func (camera *Camera) GetOwner() *base.Entity {
	return camera.owner
}

func (camera *Camera) ModelMatrix() mgl32.Mat4 {
	return mgl32.Ident4()
}

func (camera *Camera) ViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(
		camera.owner.GetTransformComponent().Position,
		camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Direction),
		camera.currentCameraComponent.Up)
}

func (camera *Camera) ProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(70, 640/480, 0.1, 16384)
}
