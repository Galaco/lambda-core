package mesh

import (
	"github.com/galaco/Lambda-Core/core/material"
	"github.com/galaco/Lambda-Core/core/texture"
)

// Generic Mesh object
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IMesh interface {
	AddVertex(...float32)
	AddNormal(...float32)
	AddUV(...float32)
	AddLightmapCoordinate(...float32)
	GenerateTangents()

	Vertices() []float32
	Normals() []float32
	UVs() []float32
	Tangents() []float32
	LightmapCoordinates() []float32

	GetMaterial() material.IMaterial
	SetMaterial(material.IMaterial)
	GetLightmap() texture.ITexture
	SetLightmap(texture.ITexture)

	//Bind()
	//Destroy()
}
