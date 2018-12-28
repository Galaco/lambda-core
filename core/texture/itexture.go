package texture

import "github.com/galaco/gosigl"

// ITexture Interface for a GPU texture
type ITexture interface {
	Bind()
	Width() int
	Height() int
	Format() gosigl.PixelFormat
	PixelDataForFrame(int) []byte
	Finish()
}
