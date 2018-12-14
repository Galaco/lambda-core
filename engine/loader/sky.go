package loader

import (
	material2 "github.com/galaco/Gource-Engine/engine/loader/material"
	"github.com/galaco/Gource-Engine/engine/texture"
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
	"github.com/galaco/Gource-Engine/engine/model"
)

// LoadSky loads the skymaterial cubemap.
// The materialname is normally obtained from the worldspawn entity
func LoadSky(materialName string) *model.Model {
	sky := model.NewModel(materialName)

	mats := make([]texture.ITexture, 6)

	mats[0] = material2.LoadSingleVtf(materialName + "up")
	mats[1] = material2.LoadSingleVtf(materialName + "dn")
	mats[2] = material2.LoadSingleVtf(materialName + "lf")
	mats[3] = material2.LoadSingleVtf(materialName + "rt")
	mats[4] = material2.LoadSingleVtf(materialName + "ft")
	mats[5] = material2.LoadSingleVtf(materialName + "bk")

	sky.AddMesh(primitive.NewCube())

	sky.GetMeshes()[0].SetMaterial(texture.NewCubemap(mats))

	return sky
}
