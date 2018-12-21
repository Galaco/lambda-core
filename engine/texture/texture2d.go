package texture

import (
	"github.com/galaco/vtf"
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Generic GPU material struct
type Texture2D struct {
	filePath string
	Buffer   uint32
	width    int
	height   int
	vtf      *vtf.Vtf
}

// Bind this material to the GPU
func (tex *Texture2D) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex.Buffer)
}

// GetFilePath Get the filepath this data was loaded from
func (tex *Texture2D) GetFilePath() string {
	return tex.filePath
}

// Width returns materials width
func (tex *Texture2D) Width() int {
	return tex.width
}

// Height returns materials height
func (tex *Texture2D) Height() int {
	return tex.height
}

// Format returns this materials colour format
func (tex *Texture2D) Format() uint32 {
	return tex.vtf.GetHeader().HighResImageFormat
}

// PixelDataForFrame get raw colour data for this frame
func (tex *Texture2D) PixelDataForFrame(frame int) []byte {
	return tex.vtf.GetHighestResolutionImageForFrame(frame)
}

// Finish Generate the GPU buffer for this material
func (tex *Texture2D) Finish() {
	gl.GenTextures(1, &tex.Buffer)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, tex.Buffer)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	if isTextureCompressed(tex.vtf.GetHeader().HighResImageFormat) {
		gl.CompressedTexImage2D(
			gl.TEXTURE_2D,
			0,
			getGLTextureFormat(tex.vtf.GetHeader().HighResImageFormat),
			int32(tex.vtf.GetHeader().Width),
			int32(tex.vtf.GetHeader().Height),
			0,
			int32(len(tex.vtf.GetHighestResolutionImageForFrame(0))),
			gl.Ptr(tex.vtf.GetHighestResolutionImageForFrame(0)))
	} else {
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RGBA,
			int32(tex.vtf.GetHeader().Width),
			int32(tex.vtf.GetHeader().Height),
			0,
			getGLTextureFormat(tex.vtf.GetHeader().HighResImageFormat),
			gl.UNSIGNED_BYTE,
			gl.Ptr(tex.vtf.GetHighestResolutionImageForFrame(0)))
	}
}

func (tex *Texture2D) Destroy() {

}

// NewMaterial returns a new material from Vtf
func NewTexture2D(filePath string, vtf *vtf.Vtf, width int, height int) *Texture2D {
	return &Texture2D{
		filePath: filePath,
		width:    width,
		height:   height,
		vtf:      vtf,
	}
}

// isTextureCompressed is a simple check for raw colour format compression
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

// getGLTextureFormat swap vtf format to openGL format
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
