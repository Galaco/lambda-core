package material

import (
	"github.com/galaco/bsp/primitives/common"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Lightmap struct {
	Color
}

// Bind this material to the GPU
func (material *Lightmap) Bind() {
	gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, material.Buffer)
}

func (material *Lightmap) Finish() {
	gl.GenTextures(1, &material.Buffer)

	material.bindInternal(gl.TEXTURE1)
}

func (material *Lightmap) bindInternal(textureSlot uint32) {
	gl.GenTextures(1, &material.Buffer)
	gl.ActiveTexture(textureSlot)
	gl.BindTexture(gl.TEXTURE_2D, material.Buffer)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(material.width),
		int32(material.height),
		0,
		material.Format(),
		gl.UNSIGNED_BYTE,
		gl.Ptr(material.rawColourData))
}

func LightmapFromColorRGBExp32(width int, height int, colorMaps []common.ColorRGBExponent32) *Lightmap {
	raw := make([]uint8, len(colorMaps) * 3)

	for idx,sample := range colorMaps {
		raw[idx * 3] = sample.R// * sample.Exponent
		raw[idx * 3 + 1] = sample.G// * sample.Exponent
		raw[idx * 3 + 2] = sample.B// * sample.Exponent
	}

	mat := &Lightmap{}

	mat.width = width
	mat.height = height
	mat.rawColourData = raw
	mat.filePath = "__lightmap"

	return mat
}