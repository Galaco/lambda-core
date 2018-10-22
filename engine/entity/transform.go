package entity

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Represents the transformation of an entity in
// a 3-dimensional space: position, rotation and scale.
// Note: Rotation is measured in degrees
type Transform struct {
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Scale    mgl32.Vec3

	prevPosition mgl32.Vec3
	prevRotation mgl32.Vec3
	prevScale    mgl32.Vec3
	matrix       mgl32.Mat4
}

func (transform *Transform) GetTransformationMatrix() mgl32.Mat4 {
	if !transform.Position.ApproxEqual(transform.prevPosition) ||
		!transform.Rotation.ApproxEqual(transform.prevRotation) ||
		!transform.Scale.ApproxEqual(transform.prevScale) {
		// Scale of 0 is invalid
		if transform.Scale.X() == 0 ||
			transform.Scale.Y() == 0 ||
			transform.Scale.Z() == 0 {
			transform.Scale = mgl32.Vec3{1, 1, 1}
		}

		//Translate
		translation := mgl32.Translate3D(transform.Position.X(), transform.Position.Y(), transform.Position.Z())

		// rotate
		rotation := mgl32.Ident4()
		//@TODO ROTATIONS

		// scale
		scale := mgl32.Scale3D(transform.Scale.X(), transform.Scale.Y(), transform.Scale.Z())

		transform.prevPosition = transform.Position
		transform.prevRotation = transform.Rotation
		transform.prevScale = transform.Scale

		transform.matrix = translation.Mul4(rotation).Mul4(scale)
	}

	return transform.matrix
}
