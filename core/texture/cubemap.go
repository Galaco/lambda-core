package texture

import (
	"github.com/galaco/gosigl"
)

// Cubemap is a 6-sided edgeless texture that can be mapped to a cube,
// Used mainly for pre-computed reflections
type Cubemap struct {
	Texture2D
	Faces []ITexture
}

// Bind this material to the GPU
func (material *Cubemap) Bind() {
	gosigl.BindTextureCubemap(gosigl.TextureSlot(0), material.Buffer)
}

// Width Get material width.
// Must have exactly 6 faces, and all faces are assumed the same size
func (material *Cubemap) Width() int {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Width()
}

// Height Get material height.
// Must have exactly 6 faces, and all faces are assumed the same size
func (material *Cubemap) Height() int {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Height()
}

// Format get material format
// Same format for all faces assumed
func (material *Cubemap) Format() gosigl.PixelFormat {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Format()
}

// Finish Generate the GPU buffer for this material
func (material *Cubemap) Finish() {
	pixelData := [6][]byte{}
	for i := 0; i < 6; i++ {
		pixelData[i] = material.Faces[i].PixelDataForFrame(0)
	}

	firstFace := material.Faces[0]
	material.Buffer = gosigl.CreateTextureCubemap(gosigl.TextureSlot(0), firstFace.Width(), firstFace.Height(), pixelData, firstFace.Format(), true)
}

func (material *Cubemap) Destroy() {

}

// NewCubemap returns a new cubemap material
func NewCubemap(materials []ITexture) *Cubemap {
	return &Cubemap{
		Faces: materials,
	}
}
