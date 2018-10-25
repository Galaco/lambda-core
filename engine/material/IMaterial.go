package material

// Interface for a GPU texture
type IMaterial interface {
	Bind()
	Width() int
	Height() int
	Format() uint32
	PixelDataForFrame(int) []byte
}
