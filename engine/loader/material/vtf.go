package material

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/lib/vtf"
	"strings"
)

// LoadMaterialList Load all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(materialList []string) {
	loadMaterials(materialList...)
}

// loadMaterials "private" function that actually does the loading
func loadMaterials(materialList ...string) (missingList []string) {
	ResourceManager := filesystem.Manager()

	// Ensure that error texture is available
	ResourceManager.Add(material.NewError(ResourceManager.ErrorTextureName()))

	materialBasePath := "materials/"

	for _, materialPath := range materialList {
		vtfTexturePath := ""

		if !strings.HasSuffix(materialPath, ".vmt") {
			materialPath += ".vmt"
		}
		// Only load the filesystem once
		if ResourceManager.Get(materialBasePath+materialPath) == nil {
			if !readVmt(materialBasePath, materialPath) {
				debug.Warn("Unable to parse: " + materialBasePath + materialPath)
				missingList = append(missingList, materialPath)
				continue
			}
			vmt := ResourceManager.Get(materialBasePath + materialPath).(*Vmt)

			// NOTE: in patch vmts include is not supported
			if vmt.GetProperty("baseTexture").AsString() != "" {
				vtfTexturePath = vmt.GetProperty("baseTexture").AsString() + ".vtf"
			}

			if vtfTexturePath != "" && !ResourceManager.Has(vtfTexturePath) {
				if !readVtf(materialBasePath, vtfTexturePath) {
					debug.Warn("Could not find: " + materialBasePath + materialPath)
					missingList = append(missingList, vtfTexturePath)
				}
			}
		}
	}

	// @TODO
	// All missing textures should be replaced with Color texture

	return missingList
}

// LoadSingleMaterial loads a single material with known file path
func LoadSingleMaterial(filePath string) material.IMaterial {
	result := loadMaterials(filePath)
	if len(result) > 0 {
		// Color
		return filesystem.Manager().Get(filesystem.Manager().ErrorTextureName()).(material.IMaterial)
	}

	vmt := filesystem.Manager().Get("materials/" + filePath).(*Vmt)
	vtfPath := vmt.GetProperty("basetexture").AsString() + ".vtf"
	if len(vtfPath) < 11 || !filesystem.Manager().Has("materials/"+vtfPath) { // 11 because len("materials/<>")
		return filesystem.Manager().Get(filesystem.Manager().ErrorTextureName()).(material.IMaterial)
	}
	return filesystem.Manager().Get("materials/" + vtfPath).(material.IMaterial)
}

func LoadSingleVtf(filePath string) material.IMaterial {
	if !readVtf("materials/", filePath) {
		return filesystem.Manager().Get(filesystem.Manager().ErrorTextureName()).(material.IMaterial)
	}
	return filesystem.Manager().Get("materials/" + filePath).(material.IMaterial)
}

func readVmt(basePath string, filePath string) bool {
	ResourceManager := filesystem.Manager()
	path := basePath + filePath

	stream, err := filesystem.Load(path)
	if err != nil {
		return false
	}

	vmt, err := ParseVmt(path, stream)
	if err != nil {
		debug.Error(err)
		return false
	}
	// Add filesystem
	ResourceManager.Add(vmt)
	return true
}

func readVtf(basePath string, filePath string) bool {
	ResourceManager := filesystem.Manager()
	stream, err := filesystem.Load(basePath + filePath)
	if err != nil {
		return false
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	texture, err := vtf.ReadFromStream(stream)
	if err != nil {
		debug.Error(err)
		return false
	}
	// Store filesystem containing raw data in memory
	ResourceManager.Add(
		material.NewMaterial(
			basePath+filePath,
			texture,
			int(texture.GetHeader().Width),
			int(texture.GetHeader().Height)))

	// Finally generate the gpu buffer for the material
	ResourceManager.Get(basePath + filePath).(material.IMaterial).Finish()
	return true
}
