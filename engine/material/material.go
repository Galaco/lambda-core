package material

import (
	"github.com/galaco/vtf"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Generic GPU material struct
type Material struct {
	filePath      string
	Buffer        uint32
	width         int
	height        int
	rawColourData []uint8
	vtf           *vtf.Vtf
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

func (material *Material) Width() int {
	return material.width
}

func (material *Material) Height() int {
	return material.height
}

func (material *Material) Format() uint32 {
	return material.vtf.GetHeader().HighResImageFormat
}

func (material *Material) PixelDataForFrame(frame int) []byte {
	return material.vtf.GetHighestResolutionImageForFrame(frame)
}

// Generate the GPU buffer for this material
func (material *Material) Finish() {
	gl.GenTextures(1, &material.Buffer)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, material.Buffer)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	if isTextureCompressed(material.vtf.GetHeader().HighResImageFormat) {
		gl.CompressedTexImage2D(
			gl.TEXTURE_2D,
			0,
			getGLTextureFormat(material.vtf.GetHeader().HighResImageFormat),
			int32(material.vtf.GetHeader().Width),
			int32(material.vtf.GetHeader().Height),
			0,
			int32(len(material.vtf.GetHighestResolutionImageForFrame(0))),
			gl.Ptr(material.vtf.GetHighestResolutionImageForFrame(0)))
	} else {
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(material.vtf.GetHeader().Width),
			int32(material.vtf.GetHeader().Height),
			0,
			getGLTextureFormat(material.vtf.GetHeader().HighResImageFormat),
			gl.UNSIGNED_BYTE,
			gl.Ptr(material.vtf.GetHighestResolutionImageForFrame(0)))
	}
}

func NewMaterial(filePath string, vtf *vtf.Vtf, width int, height int) *Material {
	return &Material{
		filePath: filePath,
		width:    width,
		height:   height,
		vtf:      vtf,
	}
}

func isTextureCompressed(vtfFormat uint32) bool {
	switch vtfFormat {
	case 13:
		return true
	case 14:
		return true
	case 15:
		return true
	}

	return false
}

func getGLTextureFormat(vtfFormat uint32) uint32 {
	switch vtfFormat {
	case 0:
		return gl.RGBA
	case 2:
		return gl.RGB
	case 3:
		return gl.BGR
	case 12:
		return gl.BGRA
	case 13:
		return gl.COMPRESSED_RGBA_S3TC_DXT1_EXT
	case 14:
		return gl.COMPRESSED_RGBA_S3TC_DXT3_EXT
	case 15:
		return gl.COMPRESSED_RGBA_S3TC_DXT5_EXT
	default:
		return gl.RGB
	}
}
