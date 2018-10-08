package bsp

import (
	"errors"
	"github.com/galaco/Gource/components"
	"github.com/galaco/Gource/components/renderable"
	"github.com/galaco/Gource/components/renderable/material"
	"github.com/galaco/Gource/engine/base/primitive"
	"github.com/galaco/Gource/engine/filesystem"
	"github.com/galaco/Gource/engine/interfaces"
	material2 "github.com/galaco/Gource/valve/loaders/material"
)

var skySuffixes = [6]string{
	"lf",
	"rt",
	"ft",
	"bk",
	"up",
	"dn",
}

func LoadSky(skyName string) (*components.Skybox, error) {
	material2.LoadSkyboxTextures(skyName)

	mats := []interfaces.IFile{
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[0] + ".vtf"),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[1] + ".vtf"),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[2] + ".vtf"),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[3] + ".vtf"),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[4] + ".vtf"),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[5] + ".vtf"),
	}

	for _, mat := range mats {
		if mat == nil {
			return nil, errors.New("failed to load cubemap: " + skyName)
		}
	}

	cubeMap := material.NewCubemap([]*material.Material{
		mats[0].(*material.Material),
		mats[1].(*material.Material),
		mats[2].(*material.Material),
		mats[3].(*material.Material),
		mats[4].(*material.Material),
		mats[5].(*material.Material),
	})
	cubeMap.GenerateGPUBuffer()

	sky := components.NewSkybox(cubeMap)

	cube := primitive.NewCube()
	cube.AddMaterial(cubeMap)
	resource := renderable.NewGPUResource([]interfaces.IPrimitive{cube})
	resource.Prepare()
	sky.AddRenderableResource(resource)

	return sky, nil
}
