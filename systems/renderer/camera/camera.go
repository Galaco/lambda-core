package camera

import (
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/go-gl/mathgl/mgl32"
	"log"
)

type Camera struct {
	base.Entity
	currentCameraComponent *components.CameraComponent
	owner *base.Entity
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

	log.Println(camera.owner.GetTransformComponent().Position)
}

func (camera *Camera) ModelMatrix() mgl32.Mat4 {
	return mgl32.Ident4()
}

func (camera *Camera) ViewMatrix() mgl32.Mat4 {
	return mgl32.LookAt(
		camera.owner.GetTransformComponent().Position.X(),
		camera.owner.GetTransformComponent().Position.Y(),
		camera.owner.GetTransformComponent().Position.Z(),
		camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Up).X(),
		camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Up).Y(),
		camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Up).Z(),
		camera.currentCameraComponent.Up.X(),
		camera.currentCameraComponent.Up.Y(),
		camera.currentCameraComponent.Up.Z())

	//return mgl32.LookAt(
	//	3,//camera.owner.GetTransformComponent().Position.X(),
	//	3,//camera.owner.GetTransformComponent().Position.Y(),
	//	3,//camera.owner.GetTransformComponent().Position.Z(),
	//	0,//camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Up).X(),
	//	0,//camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Up).Y(),
	//	0,//camera.owner.GetTransformComponent().Position.Add(camera.currentCameraComponent.Up).Z(),
	//	0,//camera.currentCameraComponent.Up.X(),
	//	1,//camera.currentCameraComponent.Up.Y(),
	//	0)//camera.currentCameraComponent.Up.Z())
}

func (camera *Camera) ProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(360), 640/480, 0.1, 100)
}