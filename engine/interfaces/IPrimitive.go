package interfaces

// Generic renderable primitive.
// Isn't necessarily a primitive, e.g. may be a submesh of a larger object
type IPrimitive interface {
	GetVertices() []float32
	GetIndices() []uint16
	GetNormals() []float32
	GetTextureCoordinates() []float32
	GetMaterial() IMaterial
	GetFaceMode() uint32
	GenerateGPUBuffer()
	Bind()
}
