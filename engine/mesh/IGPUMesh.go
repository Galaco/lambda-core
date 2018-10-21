package mesh

import "github.com/galaco/Gource-Engine/engine/mesh/primitive"

// Generic Mesh object
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IGPUMesh interface {
	Prepare()
	AddPrimitive(primitive primitive.IPrimitive)
	AddPrimitives(primitives []primitive.IPrimitive)
	GetPrimitives() []primitive.IPrimitive
}
