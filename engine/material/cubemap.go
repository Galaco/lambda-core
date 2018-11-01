package material

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

var cubeMapImageType = [6]uint32{
	gl.TEXTURE_CUBE_MAP_POSITIVE_X,
	gl.TEXTURE_CUBE_MAP_NEGATIVE_X,
	gl.TEXTURE_CUBE_MAP_POSITIVE_Y,
	gl.TEXTURE_CUBE_MAP_NEGATIVE_Y,
	gl.TEXTURE_CUBE_MAP_POSITIVE_Z,
	gl.TEXTURE_CUBE_MAP_NEGATIVE_Z,
}

type Cubemap struct {
	Material
	Faces []IMaterial
}

// Bind this material to the GPU
func (material *Cubemap) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, material.Buffer)
}

func (material *Cubemap) Width() int {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Width()
}

func (material *Cubemap) Height() int {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Height()
}

func (material *Cubemap) Format() uint32 {
	if len(material.Faces) != 6 {
		return 0
	}
	return material.Faces[0].Format()
}

// Generate the GPU buffer for this material
func (material *Cubemap) Finish() {
	gl.GenTextures(1, &material.Buffer)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_CUBE_MAP, material.Buffer)

	for i := 0; i < 6; i++ {
		cubeFace := material.Faces[i]
		gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_CUBE_MAP, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

		if isTextureCompressed(cubeFace.Format()) {
			gl.CompressedTexImage2D(
				cubeMapImageType[i],
				0,
				getGLTextureFormat(cubeFace.Format()),
				int32(cubeFace.Width()),
				int32(cubeFace.Height()),
				0,
				int32(len(cubeFace.PixelDataForFrame(0))),
				gl.Ptr(cubeFace.PixelDataForFrame(0)))
		} else {
			gl.TexImage2D(
				cubeMapImageType[i],
				0,
				gl.RGBA,
				int32(cubeFace.Width()),
				int32(cubeFace.Height()),
				0,
				getGLTextureFormat(cubeFace.Format()),
				gl.UNSIGNED_BYTE,
				gl.Ptr(cubeFace.PixelDataForFrame(0)))
		}
	}
}

func NewCubemap(materials []IMaterial) *Cubemap {
	return &Cubemap{
		Faces: materials,
	}
}
