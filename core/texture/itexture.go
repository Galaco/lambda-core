package texture

// ITexture Interface for a GPU texture
type ITexture interface {
	Bind()
	Width() int
	Height() int
	Format() uint32
	PixelDataForFrame(int) []byte
	Finish()
}
