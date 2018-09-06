package bsp

import (
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/components/renderable"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/base/primitive"
	"github.com/galaco/go-me-engine/engine/filesystem"
	"github.com/galaco/go-me-engine/engine/interfaces"
	material2 "github.com/galaco/go-me-engine/valve/loaders/material"
)

var skySuffixes = [6]string{
	"lf",
	"rt",
	"ft",
	"bk",
	"up",
	"dn",
}

func LoadSky(skyName string) *components.Skybox {
	material2.LoadSkyboxTextures(skyName)

	cubeMap := material.NewCubemap([]*material.Material{
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[0] + ".vtf").(*material.Material),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[1] + ".vtf").(*material.Material),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[2] + ".vtf").(*material.Material),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[3] + ".vtf").(*material.Material),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[4] + ".vtf").(*material.Material),
		filesystem.GetFileManager().GetFile(skyName + skySuffixes[5] + ".vtf").(*material.Material),
	})
	cubeMap.GenerateGPUBuffer()

	sky := components.NewSkybox(cubeMap)

	cube := primitive.NewCube()
	cube.AddMaterial(cubeMap)
	resource := renderable.NewGPUResource([]interfaces.IPrimitive{cube})
	resource.Prepare()
	sky.AddRenderableResource(resource)

	return sky
}
