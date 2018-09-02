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

	prevPosition mgl32.Vec3
	prevRotation mgl32.Vec3
	prevScale    mgl32.Vec3
	matrix       mgl32.Mat4
}

func (component *TransformComponent) GetTransformationMatrix() mgl32.Mat4 {
	if !component.Position.ApproxEqual(component.prevPosition) ||
		!component.Rotation.ApproxEqual(component.prevRotation) ||
		!component.Scale.ApproxEqual(component.prevScale) {
		// Scale of 0 is invalid
		if component.Scale.X() == 0 ||
			component.Scale.Y() == 0 ||
			component.Scale.Z() == 0 {
			component.Scale = mgl32.Vec3{1, 1, 1}
		}

		//Translate
		translation := mgl32.Translate3D(component.Position.X(), component.Position.Y(), component.Position.Z())

		// rotate
		rotation := mgl32.Ident4()
		//@TODO ROTATIONS

		// scale
		scale := mgl32.Scale3D(component.Scale.X(), component.Scale.Y(), component.Scale.Z())

		component.prevPosition = component.Position
		component.prevRotation = component.Rotation
		component.prevScale = component.Scale

		component.matrix = translation.Mul4(rotation).Mul4(scale)
	}

	return component.matrix
}
