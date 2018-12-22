package texture

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Colour2D is a material defined by raw/computed colour data,
// rather than loaded vtf data
type Colour2D struct {
	Texture2D
	rawColourData []uint8
}

// Format returns colour format
func (error *Colour2D) Format() uint32 {
	return gl.RGB
}

// PixelDataForFrame returns raw colour data for specific animation
// frame
func (error *Colour2D) PixelDataForFrame(frame int) []byte {
	return error.rawColourData
}

// Finish binds colour data to GPU
func (error *Colour2D) Finish() {
	gl.GenTextures(1, &error.Buffer)

	error.bindInternal(gl.TEXTURE0)
}

// bindInternal provides calls to openGL to bind colour data to GPU
func (error *Colour2D) bindInternal(textureSlot uint32) {
	gl.GenTextures(1, &error.Buffer)
	gl.ActiveTexture(textureSlot)
	gl.BindTexture(gl.TEXTURE_2D, error.Buffer)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(error.width),
		int32(error.height),
		0,
		error.Format(),
		gl.UNSIGNED_BYTE,
		gl.Ptr(error.rawColourData))
}

// Get New Error material
func NewError(name string) *Colour2D {
	mat := Colour2D{}

	mat.width = 8
	mat.height = 8
	mat.filePath = name

	// This generates purple & black chequers.
	mat.rawColourData = []uint8{
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
	}

	mat.Finish()

	return &mat
}
