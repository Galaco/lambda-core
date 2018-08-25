package interfaces

type IGPUMesh interface {
	Prepare()
	AddPrimitive(primitive IPrimitive)
	AddPrimitives(primitives []IPrimitive)
	GetPrimitives() []IPrimitive
}