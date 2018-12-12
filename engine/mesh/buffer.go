package mesh

type buffer struct {
	Vbo           uint32
	Vao           uint32
	NormalBuffer  uint32
	IndicesBuffer uint32
	UvBuffer      uint32
	LightmapUvBuffer      uint32

	FaceMode   uint32
	IsPrepared bool
}
