package material

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Generic GPU material struct
type Material struct {
	filePath string
	Buffer   uint32
	width    int
	height   int
	rawColourData []uint8
}

// Bind this material to the GPU
func (material *Material) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, material.Buffer)
}

// Get the filepath this data was loaded from
func (material *Material) GetFilePath() string {
	return material.filePath
}

func (material *Material) GetWidth() int {
	return material.width
}

func (material *Material) GetHeight() int {
	return material.height
}

// Generate the GPU buffer for this material
func (material *Material) GenerateGPUBuffer() {
	gl.GenTextures(1, &material.Buffer)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, material.Buffer)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.BGRA,
		int32(material.width),
		int32(material.height),
		0,
		gl.RGB,
		gl.UNSIGNED_BYTE,
		gl.Ptr(material.rawColourData))
}

func NewMaterial(filepath string, width int, height int, rgbData []uint8) *Material {
	return &Material{
		filePath: filepath,
		rawColourData: rgbData,
		width:    width,
		height:   height,
	}
}