package base

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Represents the transformation of an entity in
// a 3-dimensional space: position, rotation and scale.
// Note: Rotation is measured in degrees
type TransformComponent struct {
	Component
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3
}
