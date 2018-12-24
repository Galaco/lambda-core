package loader

import (
	material2 "github.com/galaco/Gource-Engine/core/loader/material"
	"github.com/galaco/Gource-Engine/core/mesh/primitive"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/galaco/Gource-Engine/core/texture"
)

const skyboxRootDir = "skybox/"

// LoadSky loads the skymaterial cubemap.
// The materialname is normally obtained from the worldspawn entity
func LoadSky(materialName string) *model.Model {
	sky := model.NewModel(materialName)

	mats := make([]texture.ITexture, 6)

	mats[0] = material2.LoadSingleTexture(skyboxRootDir + materialName + "up")
	mats[1] = material2.LoadSingleTexture(skyboxRootDir + materialName + "dn")
	mats[2] = material2.LoadSingleTexture(skyboxRootDir + materialName + "lf")
	mats[3] = material2.LoadSingleTexture(skyboxRootDir + materialName + "rt")
	mats[4] = material2.LoadSingleTexture(skyboxRootDir + materialName + "ft")
	mats[5] = material2.LoadSingleTexture(skyboxRootDir + materialName + "bk")

	sky.AddMesh(primitive.NewCube())

	sky.GetMeshes()[0].SetMaterial(texture.NewCubemap(mats))

	return sky
}
