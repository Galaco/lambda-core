package entity

import (
	"github.com/galaco/Gource-Engine/engine/config"
	"github.com/galaco/Gource-Engine/engine/core/event/message"
	"github.com/galaco/Gource-Engine/engine/input"
	"github.com/galaco/Gource-Engine/engine/input/keyboard"
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

const cameraSpeed = float64(320)
const sensitivity = float64(0.03)

var minVerticalRotation = mgl32.DegToRad(90)
var maxVerticalRotation = mgl32.DegToRad(270)

type Camera struct {
	*Base
	Up        mgl32.Vec3
	Right     mgl32.Vec3
	Direction mgl32.Vec3
	worldUp   mgl32.Vec3
	frameTime float64
}

func (camera *Camera) ReceiveMessage(message message.IMessage) {

}

func (camera *Camera) Update(dt float64) {
	vel := cameraSpeed * dt
	if input.GetKeyboard().IsKeyDown(keyboard.KeyW) {
		camera.Transform().Position = camera.Transform().Position.Add(camera.Direction.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyA) {
		camera.Transform().Position = camera.Transform().Position.Sub(camera.Right.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyS) {
		camera.Transform().Position = camera.Transform().Position.Sub(camera.Direction.Mul(float32(vel)))
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyD) {
		camera.Transform().Position = camera.Transform().Position.Add(camera.Right.Mul(float32(vel)))
	}

	camera.Transform().Rotation[0] -= float32(input.GetMouse().GetCoordinates()[0] * sensitivity)
	camera.Transform().Rotation[2] -= float32(input.GetMouse().GetCoordinates()[1] * sensitivity)

	// Lock vertical rotation
	if camera.Transform().Rotation[2] > maxVerticalRotation {
		camera.Transform().Rotation[2] = maxVerticalRotation
	}
	if camera.Transform().Rotation[2] < minVerticalRotation {
		camera.Transform().Rotation[2] = minVerticalRotation
	}

	camera.updateVectors()
}

// Update the camera directional properties with any changes
func (camera *Camera) updateVectors() {
	rot := camera.Transform().Rotation

	// Calculate the new Front vector
	camera.Direction = mgl32.Vec3{
		float32(math.Cos(float64(rot[2])) * math.Sin(float64(rot[0]))),
		float32(math.Cos(float64(rot[2])) * math.Cos(float64(rot[0]))),
		float32(math.Sin(float64(rot[2]))),
	}
	// Also re-calculate the Right and Up vector
	camera.Right = mgl32.Vec3{
		float32(math.Sin(float64(rot[0]) - math.Pi/2)),
		float32(math.Cos(float64(rot[0]) - math.Pi/2)),
		0,
	}
	camera.Up = camera.Right.Cross(camera.Direction)
}

func (camera *Camera) ModelMatrix() mgl32.Mat4 {
	return mgl32.Ident4()
}

func (camera *Camera) ViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(
		camera.Transform().Position,
		camera.Transform().Position.Add(camera.Direction),
		camera.Up)
}

func (camera *Camera) ProjectionMatrix() mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(70), float32(config.Get().Video.Width)/float32(config.Get().Video.Height), 0.1, 16384)
}

func NewCamera() *Camera {
	return &Camera{
		Base:      &Base{},
		Up:        mgl32.Vec3{0, 1, 0},
		worldUp:   mgl32.Vec3{0, 1, 0},
		Direction: mgl32.Vec3{0, 0, -1},
	}
}
