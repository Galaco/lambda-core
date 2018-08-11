package entity

import (
	"github.com/go-gl/mathgl/mgl32"
)

type TransformComponent struct {
	Component
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale mgl32.Vec3
}