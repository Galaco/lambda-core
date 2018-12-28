package texture

import "github.com/galaco/Gource-Engine/glapi"

// ITexture Interface for a GPU texture
type ITexture interface {
	Bind()
	Width() int
	Height() int
	Format() glapi.PixelFormat
	PixelDataForFrame(int) []byte
	Finish()
}
