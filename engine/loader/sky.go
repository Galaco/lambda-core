package loader

import (
	material2 "github.com/galaco/Gource-Engine/engine/loader/material"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
	"github.com/galaco/Gource-Engine/engine/model"
)

func LoadSky(materialName string) *model.Model {
	sky := model.NewModel(materialName)

	mats := make([]material.IMaterial, 6)

	mats[0] = material2.LoadSingleVtf(materialName + "up")
	mats[1] = material2.LoadSingleVtf(materialName + "dn")
	mats[2] = material2.LoadSingleVtf(materialName + "lf")
	mats[3] = material2.LoadSingleVtf(materialName + "rt")
	mats[4] = material2.LoadSingleVtf(materialName + "ft")
	mats[5] = material2.LoadSingleVtf(materialName + "bk")

	sky.AddMesh(primitive.NewCube())

	sky.GetMeshes()[0].SetMaterial(material.NewCubemap(mats))

	return sky
}
