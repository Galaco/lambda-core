package camera

import (
	"github.com/galaco/Gource-Engine/components"
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/core/event/message"
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/factory"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messages"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messagetype"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	currentCameraComponent *components.CameraComponent
	owner                  *entity.Base
}

func (camera *Camera) Initialize() {
	event.GetEventManager().Listen(messagetype.ChangeActiveCamera, camera)
}

func (camera *Camera) ReceiveMessage(message message.IMessage) {
	if message.GetType() == messagetype.ChangeActiveCamera {
		camera.currentCameraComponent = message.(*messages.ChangeActiveCamera).Component.(*components.CameraComponent)
		camera.owner = factory.GetObjectManager().GetEntityByHandle(camera.currentCameraComponent.GetOwnerHandle()).(*entity.Base)
	}
}

func (camera *Camera) SendMessage() message.IMessage {
	return nil
}

func (camera *Camera) Update(dt float64) {
	camera.currentCameraComponent.Update(dt)
}

func (camera *Camera) GetOwner() *entity.Base {
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
