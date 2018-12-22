package mesh

import (
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/texture"
)

// Generic Mesh object
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IMesh interface {
	AddVertex(...float32)
	AddNormal(...float32)
	AddTextureCoordinate(...float32)
	AddLightmapCoordinate(...float32)
	Finish()

	Vertices() []float32
	Normals() []float32
	TextureCoordinates() []float32
	LightmapCoordinates() []float32

	GetMaterial() material.IMaterial
	SetMaterial(material.IMaterial)
	GetLightmap() texture.ITexture
	SetLightmap(texture.ITexture)

	Bind()
	Destroy()
}
