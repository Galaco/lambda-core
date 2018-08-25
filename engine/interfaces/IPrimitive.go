package interfaces

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
