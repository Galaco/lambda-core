package loader

import (
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
	"github.com/galaco/Gource-Engine/engine/model"
)

func LoadSky(materialName string) *model.Model {
	sky := model.NewModel(materialName)

	mats := make([]material.IMaterial, 6)

	mats[0] = material.LoadSingleVtf(materialName + "up")
	mats[1] = material.LoadSingleVtf(materialName + "dn")
	mats[2] = material.LoadSingleVtf(materialName + "lf")
	mats[3] = material.LoadSingleVtf(materialName + "rt")
	mats[4] = material.LoadSingleVtf(materialName + "ft")
	mats[5] = material.LoadSingleVtf(materialName + "bk")

	sky.AddMesh(primitive.NewCube())

	sky.GetMeshes()[0].SetMaterial(material.NewCubemap(mats))

	return sky
}
