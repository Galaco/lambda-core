package glapi

import (
	opengl "github.com/go-gl/gl/v4.1-core/gl"
)

type TextureBindingId uint32
type TextureSlotId uint32
type PixelFormat uint32

const RGBA = PixelFormat(opengl.RGBA)
const RGB = PixelFormat(opengl.RGB)
const BGR = PixelFormat(opengl.BGR)
const BGRA = PixelFormat(opengl.BGRA)
const DXT1 = PixelFormat(opengl.COMPRESSED_RGB_S3TC_DXT1_EXT)
const DXT1A = PixelFormat(opengl.COMPRESSED_RGBA_S3TC_DXT1_EXT)
const DXT3 = PixelFormat(opengl.COMPRESSED_RGBA_S3TC_DXT3_EXT)
const DXT5 = PixelFormat(opengl.COMPRESSED_RGBA_S3TC_DXT5_EXT)

var textureSlots = [...]TextureSlotId{
	opengl.TEXTURE0,
	opengl.TEXTURE1,
	opengl.TEXTURE2,
	opengl.TEXTURE3,
	opengl.TEXTURE4,
	opengl.TEXTURE5,
	opengl.TEXTURE6,
	opengl.TEXTURE7,
	opengl.TEXTURE8,
	opengl.TEXTURE9,
	opengl.TEXTURE10,
	opengl.TEXTURE11,
	opengl.TEXTURE12,
	opengl.TEXTURE13,
	opengl.TEXTURE14,
	opengl.TEXTURE15,
	opengl.TEXTURE16,
	opengl.TEXTURE17,
	opengl.TEXTURE18,
	opengl.TEXTURE19,
	opengl.TEXTURE20,
	opengl.TEXTURE21,
	opengl.TEXTURE22,
	opengl.TEXTURE23,
	opengl.TEXTURE24,
	opengl.TEXTURE25,
	opengl.TEXTURE26,
	opengl.TEXTURE27,
	opengl.TEXTURE28,
	opengl.TEXTURE29,
	opengl.TEXTURE30,
	opengl.TEXTURE31,
}

var cubemapSlots = [...]TextureSlotId{
	opengl.TEXTURE_CUBE_MAP_POSITIVE_X,
	opengl.TEXTURE_CUBE_MAP_NEGATIVE_X,
	opengl.TEXTURE_CUBE_MAP_POSITIVE_Y,
	opengl.TEXTURE_CUBE_MAP_NEGATIVE_Y,
	opengl.TEXTURE_CUBE_MAP_POSITIVE_Z,
	opengl.TEXTURE_CUBE_MAP_NEGATIVE_Z,
}

func TextureSlot(index int) TextureSlotId {
	if index > 31 {
		return textureSlots[0]
	}
	return textureSlots[index]
}

func BindTexture2D(slot TextureSlotId, id TextureBindingId) {
	opengl.ActiveTexture(uint32(slot))
	opengl.BindTexture(opengl.TEXTURE_2D, uint32(id))
}

func BindTextureCubemap(slot TextureSlotId, id TextureBindingId) {
	opengl.ActiveTexture(uint32(slot))
	opengl.BindTexture(opengl.TEXTURE_CUBE_MAP, uint32(id))
}

func CreateTexture2D(slot TextureSlotId, width int, height int, pixelData []byte, pixelFormat PixelFormat, clampToEdge bool) TextureBindingId {
	textureBuffer := uint32(0)
	opengl.GenTextures(1, &textureBuffer)
	BindTexture2D(slot, TextureBindingId(textureBuffer))

	createTexture(opengl.TEXTURE_2D, opengl.TEXTURE_2D, width, height, pixelData, pixelFormat, clampToEdge)

	return TextureBindingId(textureBuffer)
}

func CreateTextureCubemap(slot TextureSlotId, width int, height int, pixelData [6][]byte, pixelFormat PixelFormat, clampToEdge bool) TextureBindingId {
	textureBuffer := uint32(0)
	opengl.GenTextures(1, &textureBuffer)
	BindTextureCubemap(slot, TextureBindingId(textureBuffer))

	createTexture(opengl.TEXTURE_CUBE_MAP, cubemapSlots[0], width, height, pixelData[0], pixelFormat, clampToEdge)
	createTexture(opengl.TEXTURE_CUBE_MAP, cubemapSlots[1], width, height, pixelData[1], pixelFormat, clampToEdge)
	createTexture(opengl.TEXTURE_CUBE_MAP, cubemapSlots[2], width, height, pixelData[2], pixelFormat, clampToEdge)
	createTexture(opengl.TEXTURE_CUBE_MAP, cubemapSlots[3], width, height, pixelData[3], pixelFormat, clampToEdge)
	createTexture(opengl.TEXTURE_CUBE_MAP, cubemapSlots[4], width, height, pixelData[4], pixelFormat, clampToEdge)
	createTexture(opengl.TEXTURE_CUBE_MAP, cubemapSlots[5], width, height, pixelData[5], pixelFormat, clampToEdge)

	return TextureBindingId(textureBuffer)
}

func createTexture(textureType uint32, target TextureSlotId, width int, height int, pixelData []byte, pixelFormat PixelFormat, clampToEdge bool) {
	opengl.TexParameteri(textureType, opengl.TEXTURE_MIN_FILTER, opengl.LINEAR)
	opengl.TexParameteri(textureType, opengl.TEXTURE_MAG_FILTER, opengl.LINEAR)

	if clampToEdge == true {
		opengl.TexParameteri(textureType, opengl.TEXTURE_WRAP_S, opengl.CLAMP_TO_EDGE)
		opengl.TexParameteri(textureType, opengl.TEXTURE_WRAP_T, opengl.CLAMP_TO_EDGE)
	} else {
		opengl.TexParameteri(textureType, opengl.TEXTURE_WRAP_S, opengl.REPEAT)
		opengl.TexParameteri(textureType, opengl.TEXTURE_WRAP_T, opengl.REPEAT)
	}

	if isPixelFormatCompressed(pixelFormat) {
		opengl.CompressedTexImage2D(
			uint32(target),
			0,
			uint32(pixelFormat),
			int32(width),
			int32(height),
			0,
			int32(len(pixelData)),
			opengl.Ptr(pixelData))
	} else {
		opengl.TexImage2D(
			uint32(target),
			0,
			opengl.RGBA,
			int32(width),
			int32(height),
			0,
			uint32(pixelFormat),
			opengl.UNSIGNED_BYTE,
			opengl.Ptr(pixelData))
	}
}

func DeleteTextures(ids ...TextureBindingId) {
	rawIds := make([]uint32, len(ids))
	for idx, id := range ids {
		rawIds[idx] = uint32(id)
	}
	opengl.DeleteTextures(int32(len(rawIds)), &rawIds[0])
}

func isPixelFormatCompressed(format PixelFormat) bool {
	switch format {
	case DXT1, DXT1A, DXT3, DXT5:
		return true
	}

	return false
}