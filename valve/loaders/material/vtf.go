package material

import (
	"github.com/galaco/Gource-Engine/components/renderable/material"
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/valve/file"
	"github.com/galaco/Gource-Engine/valve/libwrapper/vtf"
)

// Load all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(materialList []string) {
	read(materialList)
}

func read(materialList []string) (missingList []string) {
	FileManager := filesystem.GetFileManager()
	materialBasePath := "materials/"

	for _, materialPath := range materialList {
		vtfTexturePath := ""
		// Only load the file once
		if FileManager.GetFile(materialBasePath+materialPath) == nil {
			if !readVmt(materialBasePath, materialPath) {
				debug.Log("Could not find: " + materialPath)
				missingList = append(missingList, materialPath)
				continue
			}
			vmt := FileManager.GetFile(materialPath).(*Vmt)

			// NOTE: in patch vmts include is not supported
			if vmt.GetProperty("baseTexture").AsString() != "" {
				vtfTexturePath = vmt.GetProperty("baseTexture").AsString() + ".vtf"
			}

			if vtfTexturePath != "" && FileManager.GetFile(vtfTexturePath) == nil {
				if !readVtf(materialBasePath, vtfTexturePath) {
					debug.Log("Could not find: " + materialPath)
					missingList = append(missingList, vtfTexturePath)
				}
			}
		}
	}

	return missingList
}

func readVmt(basePath string, filePath string) bool {
	FileManager := filesystem.GetFileManager()
	path := basePath + filePath + ".vmt"

	stream, err := file.Load(path)
	if err != nil {
		return false
	}

	vmt, err := ParseVmt(filePath, stream)
	if err != nil {
		debug.Log(err)
		return false
	}
	// Add file
	FileManager.AddFile(vmt)
	return true
}

func readVtf(basePath string, filePath string) bool {
	FileManager := filesystem.GetFileManager()
	stream, err := file.Load(basePath + filePath)
	if err != nil {
		return false
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	texture, err := vtf.ReadFromStream(stream)
	if err != nil {
		debug.Log(err)
		return false
	}
	// Store file containing raw data in memory
	FileManager.AddFile(
		material.NewMaterial(
			filePath,
			texture,
			int(texture.GetHeader().Width),
			int(texture.GetHeader().Height)))

	// Finally generate the gpu buffer for the material
	FileManager.GetFile(filePath).(*material.Material).GenerateGPUBuffer()
	return true
}

func LoadSkyboxTextures(skyName string) {
	exts := []string{
		"lf",
		"bk",
		"rt",
		"ft",
		"up",
		"dn",
	}

	for _, ext := range exts {
		readVtf("materials/skybox/", skyName+ext+".vtf")
	}
}
