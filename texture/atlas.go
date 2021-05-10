package texture

import (
	"github.com/galaco/lambda-core/lib/math/shape"
	"github.com/galaco/vtf/format"
)

// @TODO THIS IS NOT COMPLETE. IT DOES NOT WORK

// Atlas
// A texture atlas implementation.
type Atlas struct {
	Colour2D
}

// Format returns colour format
// For now always RGBA
func (atlas *Atlas) Format() uint32 {
	return uint32(format.RGB888)
}

// PackTextures
func (atlas *Atlas) PackTextures(textures []ITexture, padding int) ([]shape.Rect, error) {
	uvRects := make([]shape.Rect, 0)

	return uvRects, nil
}

//// findSpace finds free space in atlas buffer to write rectangle to
//func (atlas *Atlas) findSpace(width int, height int) (x, y int, err error) {
//
//	return x, y, err
//}


// NewAtlas
func NewAtlas(width int, height int) *Atlas {
	return &Atlas{
		Colour2D: Colour2D{
			rawColourData: make([]uint8, width*height*3),
			Texture2D: Texture2D{
				width:  width,
				height: height,
			},
		},
	}
}
