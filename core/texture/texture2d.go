package texture

import (
	"github.com/galaco/gosigl"
	"github.com/galaco/vtf"
)

// Generic GPU material struct
type Texture2D struct {
	filePath string
	Buffer   gosigl.TextureBindingId
	width    int
	height   int
	vtf      *vtf.Vtf
}

// Bind this material to the GPU
//func (tex *Texture2D) Bind() {
//	gosigl.BindTexture2D(gosigl.TextureSlot(0), tex.Buffer)
//}

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
func (tex *Texture2D) Format() gosigl.PixelFormat {
	return getGLTextureFormat(tex.vtf.GetHeader().HighResImageFormat)
}

// PixelDataForFrame get raw colour data for this frame
func (tex *Texture2D) PixelDataForFrame(frame int) []byte {
	return tex.vtf.GetHighestResolutionImageForFrame(frame)
}

//func (tex *Texture2D) Destroy() {
//	gosigl.DeleteTextures(tex.Buffer)
//}

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
func getGLTextureFormat(vtfFormat uint32) gosigl.PixelFormat {
	switch vtfFormat {
	case 0:
		return gosigl.RGBA
	case 2:
		return gosigl.RGB
	case 3:
		return gosigl.BGR
	case 12:
		return gosigl.BGRA
	case 13:
		return gosigl.DXT1
	case 14:
		return gosigl.DXT3
	case 15:
		return gosigl.DXT5
	default:
		return gosigl.RGB
	}
}
