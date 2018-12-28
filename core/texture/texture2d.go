package texture

import (
	"github.com/galaco/Gource-Engine/glapi"
	"github.com/galaco/vtf"
)

// Generic GPU material struct
type Texture2D struct {
	filePath string
	Buffer   glapi.TextureBindingId
	width    int
	height   int
	vtf      *vtf.Vtf
}

// Bind this material to the GPU
func (tex *Texture2D) Bind() {
	glapi.BindTexture2D(glapi.TextureSlot(0), tex.Buffer)
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
func (tex *Texture2D) Format() glapi.PixelFormat {
	return getGLTextureFormat(tex.vtf.GetHeader().HighResImageFormat)
}

// PixelDataForFrame get raw colour data for this frame
func (tex *Texture2D) PixelDataForFrame(frame int) []byte {
	return tex.vtf.GetHighestResolutionImageForFrame(frame)
}

// Finish Generate the GPU buffer for this material
func (tex *Texture2D) Finish() {
	tex.Buffer = glapi.CreateTexture2D(
		glapi.TextureSlot(0),
		int(tex.vtf.GetHeader().Width),
		int(tex.vtf.GetHeader().Height),
		tex.vtf.GetHighestResolutionImageForFrame(0),
		getGLTextureFormat(tex.vtf.GetHeader().HighResImageFormat),
		false)
}

func (tex *Texture2D) Destroy() {
	glapi.DeleteTextures(tex.Buffer)
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

// getGLTextureFormat swap vtf format to openGL format
func getGLTextureFormat(vtfFormat uint32) glapi.PixelFormat {
	switch vtfFormat {
	case 0:
		return glapi.RGBA
	case 2:
		return glapi.RGB
	case 3:
		return glapi.BGR
	case 12:
		return glapi.BGRA
	case 13:
		return glapi.DXT1
	case 14:
		return glapi.DXT3
	case 15:
		return glapi.DXT5
	default:
		return glapi.RGB
	}
}
