package material

import (
	"github.com/galaco/Gource-Engine/engine/core/logger"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/resource"
	"github.com/galaco/Gource-Engine/engine/texture"
	"github.com/galaco/Gource-Engine/lib/vtf"
	"strings"
)

const materialRootPath = "materials/"

// LoadMaterialList GetFile all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(materialList []string) {
	loadMaterials(materialList...)
}

func LoadErrorMaterial() {
	ResourceManager := resource.Manager()
	name := ResourceManager.ErrorTextureName()

	// Ensure that error texture is available
	ResourceManager.AddTexture(texture.NewError(name))
	errorMat := &material.Material{
		FilePath: name,
	}
	errorMat.Textures.BaseTexture = ResourceManager.GetTexture(name).(texture.ITexture)
	ResourceManager.AddMaterial(errorMat)
}

// loadMaterials "private" function that actually does the loading
func loadMaterials(materialList ...string) (missingList []string) {
	ResourceManager := resource.Manager()

	for _, materialPath := range materialList {
		vtfTexturePath := ""

		if !strings.HasSuffix(materialPath, ".vmt") {
			materialPath += ".vmt"
		}
		// Only load the filesystem once
		if ResourceManager.GetMaterial(materialRootPath+materialPath) == nil {
			if !readVmt(materialRootPath + materialPath) {
				logger.Warn("Unable to parse: " + materialRootPath + materialPath)
				missingList = append(missingList, materialPath)
				continue
			}
			vmt := ResourceManager.GetMaterial(materialRootPath + materialPath).(*material.Material)

			// NOTE: in patch vmts include is not supported
			if vmt.BaseTextureName != "" {
				vtfTexturePath = vmt.BaseTextureName + ".vtf"
			}

			vmt.Textures.BaseTexture = LoadSingleTexture(vtfTexturePath)
		}
	}
	return missingList
}

// LoadSingleMaterial loads a single material with known file path
func LoadSingleMaterial(filePath string) material.IMaterial {
	result := loadMaterials(filePath)
	if len(result) > 0 {
		return resource.Manager().GetMaterial(resource.Manager().ErrorTextureName()).(material.IMaterial)
	}
	return resource.Manager().GetMaterial("materials/" + filePath).(material.IMaterial)
}

func LoadSingleTexture(filePath string) texture.ITexture {
	if filePath == "" || !readVtf(materialRootPath + filePath) {
		return resource.Manager().GetTexture(resource.Manager().ErrorTextureName()).(texture.ITexture)
	}
	return resource.Manager().GetTexture(materialRootPath + filePath).(texture.ITexture)
}

func readVmt(path string) bool {
	ResourceManager := resource.Manager()

	stream, err := filesystem.GetFile(path)
	if err != nil {
		logger.Error(err)
		return false
	}

	vmt, err := ParseVmt(path, stream)
	if err != nil {
		logger.Error(err)
		return false
	}
	// Add filesystem
	mat := &material.Material{
		FilePath:        path,
		BaseTextureName: vmt.GetProperty("baseTexture").AsString(),
	}
	ResourceManager.AddMaterial(mat)
	return true
}

func readVtf(path string) bool {
	ResourceManager := resource.Manager()
	stream, err := filesystem.GetFile(path)
	if err != nil {
		return false
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	read, err := vtf.ReadFromStream(stream)
	if err != nil {
		logger.Error(err)
		return false
	}
	// Store filesystem containing raw data in memory
	ResourceManager.AddTexture(
		texture.NewTexture2D(
			path,
			read,
			int(read.GetHeader().Width),
			int(read.GetHeader().Height)))

	// Finally generate the gpu buffer for the material
	ResourceManager.GetTexture(path).(texture.ITexture).Finish()
	return true
}
