package interfaces

// Generic Mesh object
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IGPUMesh interface {
	Prepare()
	AddPrimitive(primitive IPrimitive)
	AddPrimitives(primitives []IPrimitive)
	GetPrimitives() []IPrimitive
}
