package texture

import (
	"github.com/galaco/vtf"
)

// Generic GPU material struct
type Texture2D struct {
	filePath string
	width    int
	height   int
	vtf      *vtf.Vtf
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

// Thumbnail returns a small thumbnail image of a material
func (tex *Texture2D) Thumbnail() []byte {
	return tex.vtf.GetLowResImageData()
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
