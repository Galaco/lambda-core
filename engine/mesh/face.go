package mesh

import (
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/texture"
)

type Face struct {
	offset   int32
	length   int32
	material material.IMaterial
	lightmap *texture.Lightmap
}

func (face *Face) Offset() int32 {
	return face.offset
}

func (face *Face) Length() int32 {
	return face.length
}

func (face *Face) IsLightmapped() bool {
	return face.Lightmap() != nil
}

func (face *Face) AddMaterial(mat material.IMaterial) {
	face.material = mat
}

func (face *Face) AddLightmap(lightmap *texture.Lightmap) {
	face.lightmap = lightmap
}

func (face *Face) Material() material.IMaterial {
	return face.material
}

func (face *Face) Lightmap() *texture.Lightmap {
	return face.lightmap
}

func NewFace(offset int32, length int32, mat texture.ITexture, lightmap *texture.Lightmap) Face {
	return Face{
		offset:   offset,
		length:   length,
		material: mat,
		lightmap: lightmap,
	}
}
