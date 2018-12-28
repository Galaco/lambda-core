package texture

import (
	"github.com/galaco/Gource-Engine/glapi"
	"github.com/galaco/bsp/primitives/common"
)

// Lightmap is a material used for lighting a face
type Lightmap struct {
	Colour2D
}

// Bind this material to the GPU
func (material *Lightmap) Bind() {
	glapi.BindTexture2D(glapi.TextureSlot(1), material.Buffer)
}

// Finish binds this material data to the GPU
func (material *Lightmap) Finish() {
	material.Buffer = glapi.CreateTexture2D(glapi.TextureSlot(1), material.Width(), material.Height(), material.PixelDataForFrame(0), material.Format(), true)
}
// Create a lightmap from BSP stored colour data
func LightmapFromColorRGBExp32(width int, height int, colorMaps []common.ColorRGBExponent32) *Lightmap {
	raw := make([]uint8, len(colorMaps)*3)

	for idx, sample := range colorMaps {
		raw[idx*3] = sample.R   // * sample.Exponent
		raw[idx*3+1] = sample.G // * sample.Exponent
		raw[idx*3+2] = sample.B // * sample.Exponent
	}

	mat := &Lightmap{}

	mat.width = width
	mat.height = height
	mat.rawColourData = raw
	mat.filePath = "__lightmap"

	return mat
}
