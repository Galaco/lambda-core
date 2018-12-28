package controllers

import (
	"github.com/galaco/Gource-Engine/client/input"
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/client/scene"
	"github.com/galaco/Gource-Engine/core"
)

type Camera struct {
	core.Manager
}

func (controller *Camera) Update(dt float64) {
	cam := scene.Get().CurrentCamera()
	if cam == nil {
		return
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyW) {
		cam.Forwards(dt)
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyA) {
		cam.Left(dt)
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyS) {
		cam.Backwards(dt)
	}
	if input.GetKeyboard().IsKeyDown(keyboard.KeyD) {
		cam.Right(dt)
	}

	cam.Rotate(input.GetMouse().GetCoordinates()[0], 0, input.GetMouse().GetCoordinates()[1])
}
