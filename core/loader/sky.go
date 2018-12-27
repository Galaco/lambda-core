package loader

import (
	material2 "github.com/galaco/Gource-Engine/core/loader/material"
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/mesh/primitive"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/galaco/Gource-Engine/core/texture"
)

const skyboxRootDir = "skybox/"

// LoadSky loads the skymaterial cubemap.
// The materialname is normally obtained from the worldspawn entity
func LoadSky(materialName string) *model.Model {
	sky := model.NewModel(materialName)

	mats := make([]material.IMaterial, 6)

	mats[0] = material2.LoadSingleMaterial(skyboxRootDir + materialName + "up.vmt")
	mats[1] = material2.LoadSingleMaterial(skyboxRootDir + materialName + "dn.vmt")
	mats[2] = material2.LoadSingleMaterial(skyboxRootDir + materialName + "lf.vmt")
	mats[3] = material2.LoadSingleMaterial(skyboxRootDir + materialName + "rt.vmt")
	mats[4] = material2.LoadSingleMaterial(skyboxRootDir + materialName + "ft.vmt")
	mats[5] = material2.LoadSingleMaterial(skyboxRootDir + materialName + "bk.vmt")

	texs := make([]texture.ITexture, 6)
	for i := 0; i < 6; i++ {
		texs[i] = mats[i].(*material.Material).Textures.BaseTexture
	}

	sky.AddMesh(primitive.NewCube())

	sky.GetMeshes()[0].SetMaterial(texture.NewCubemap(texs))

	return sky
}
