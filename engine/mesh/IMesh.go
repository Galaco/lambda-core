package mesh

import (
	"github.com/galaco/Gource-Engine/engine/material"
)

// Generic Mesh object
// Most renderable objects should implement this, but there
// are probably many custom cases that may not
type IMesh interface {
	AddVertex(...float32)
	AddNormal(...float32)
	AddTextureCoordinate(...float32)
	Finish()

	Vertices() []float32
 	Normals() []float32
	TextureCoordinates() []float32

	GetMaterial() material.IMaterial
	SetMaterial(material.IMaterial)

	Bind()
}
