package material

import (
	"github.com/galaco/Gource-Engine/engine/core/logger"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/resource"
	"github.com/galaco/Gource-Engine/engine/texture"
	"github.com/galaco/KeyValues"
	"strings"
)

// LoadMaterialList GetFile all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(materialList []string) {
	loadMaterials(materialList...)
}

// LoadErrorMaterial ensures that the error material has been loaded
func LoadErrorMaterial() {
	ResourceManager := resource.Manager()
	name := ResourceManager.ErrorTextureName()

	if ResourceManager.GetMaterial(name) != nil {
		return
	}

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
			if readVmt(materialRootPath + materialPath) != nil {
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
	if resource.Manager().GetMaterial(materialRootPath + filePath) != nil {
		return resource.Manager().GetMaterial(materialRootPath + filePath).(material.IMaterial)
	}
	result := loadMaterials(filePath)
	if len(result) > 0 {
		return resource.Manager().GetMaterial(resource.Manager().ErrorTextureName()).(material.IMaterial)
	}
	return resource.Manager().GetMaterial("materials/" + filePath).(material.IMaterial)
}

func readVmt(path string) error {
	ResourceManager := resource.Manager()

	stream, err := filesystem.GetFile(path)
	if err != nil {
		return err
	}

	reader := keyvalues.NewReader(stream)
	kvs,err := reader.Read()
	if err != nil {
		return err
	}
	roots,err := kvs.Children()
	if err != nil {
		return err
	}
	root := roots[0]

	baseTextureKV,err := root.Find("$basetexture")
	if err != nil {
		return err
	}
	baseTexture,err := baseTextureKV.AsString()
	if err != nil {
		return err
	}

	mat := &material.Material{
		FilePath:        path,
		BaseTextureName: baseTexture,
	}
	ResourceManager.AddMaterial(mat)
	return nil
}
