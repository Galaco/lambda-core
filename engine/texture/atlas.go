package texture

import (
	"github.com/galaco/Gource-Engine/lib/math/shape"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"sort"
)

type Atlas struct {
	Colour2D
	PackedRectangles []shape.Rect
}

// Format returns colour format
// For now always RGBA
func (atlas *Atlas) Format() uint32 {
	return gl.RGBA
}

// PackTextures
func (atlas *Atlas) PackTextures(textures []Texture2D, padding int, maxSize int) ([]shape.Rect, error) {
	// order by size smallest to largest
	ordered := make([]*Texture2D, len(textures))
	for idx := range textures {
		ordered[idx] = &textures[idx]
	}
	sort.Slice(ordered, func(i, j int) bool {
		return (ordered[i].Width() * ordered[i].Height()) > (ordered[j].Width() * ordered[j].Height())
	})

	var err error
	for _,tex := range ordered {
		x,y,err := atlas.findSpace(tex.Width(), tex.Height())
		if err != nil {
			// Atlas is too small for textures to pack
			break
		}

		// Ensure that all textures are the same format

		// Write data to buffer
		atlas.insert(tex.PixelDataForFrame(0), tex.Width(), tex.Height(), x, y)

		atlas.PackedRectangles = append(atlas.PackedRectangles, *shape.NewRect(
			mgl32.Vec2{float32(x), float32(y)},
			mgl32.Vec2{float32(x) + float32(tex.Width()), float32(y) + float32(tex.Height())}))
	}

	return atlas.PackedRectangles, err
}

// findSpace finds free space in atlas buffer to write rectangle to
func (atlas *Atlas) findSpace(width int, height int) (x, y int, err error) {



	return x,y,err
}

// insert write raw data into atlas at calculated position
func (atlas *Atlas) insert(data []uint8, width int, height int, x int, y int) {
	for i := 0; i < height; i++ {
		offset := (x * atlas.width * 4) + (y * i * 4)

		for j := 0; j < width; j++ {
			atlas.rawColourData[offset + (j * 4) + 0] = data[(i * width) + (j * 4)]
			atlas.rawColourData[offset + (j * 4) + 1] = data[(i * width) + (j * 4) + 1]
			atlas.rawColourData[offset + (j * 4) + 2] = data[(i * width) + (j * 4) + 2]
			atlas.rawColourData[offset + (j * 4) + 3] = 255
		}
	}
}

// NewAtlas
func NewAtlas(width int, height int) *Atlas {
	return &Atlas{
		Colour2D: Colour2D{
			rawColourData: make([]uint8, width * height * 4),
			Texture2D: Texture2D{
				width: width,
				height: height,
			},
		},
		PackedRectangles: make([]shape.Rect, 0),
	}
}