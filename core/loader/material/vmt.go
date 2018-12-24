package material

import (
	"errors"
	keyvalues2 "github.com/galaco/Gource-Engine/core/loader/keyvalues"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/resource"
	"github.com/galaco/Gource-Engine/core/texture"
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
			mat, err := readVmt(materialRootPath + materialPath)
			if err != nil {
				logger.Warn("Failed to load material: %s. Reason: %s", materialRootPath+materialPath, err)
				missingList = append(missingList, materialPath)
				continue
			}
			vmt := mat.(*material.Material)

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
	if resource.Manager().GetMaterial(materialRootPath+filePath) != nil {
		return resource.Manager().GetMaterial(materialRootPath + filePath).(material.IMaterial)
	}
	result := loadMaterials(filePath)
	if len(result) > 0 {
		return resource.Manager().GetMaterial(resource.Manager().ErrorTextureName()).(material.IMaterial)
	}
	return resource.Manager().GetMaterial("materials/" + filePath).(material.IMaterial)
}

func readVmt(path string) (material.IMaterial, error) {
	ResourceManager := resource.Manager()

	kvs, err := keyvalues2.ReadKeyValues(path)
	if err != nil {
		return nil, err
	}
	roots, err := kvs.Children()
	if err != nil {
		return nil, err
	}
	root := roots[0]

	include, err := root.Find("include")
	if err == nil {
		includePath, _ := include.AsString()
		root, err = mergeIncludedVmtRecursive(root, includePath)
		if err != nil {
			return nil, err
		}
	}

	// @NOTE this will be replaced with a proper kv->material builder
	baseTextureKV, err := root.Find("$basetexture")
	if err != nil {
		return nil, err
	}
	baseTexture, err := baseTextureKV.AsString()
	if err != nil {
		return nil, err
	}

	mat := &material.Material{
		FilePath:        path,
		BaseTextureName: baseTexture,
	}
	ResourceManager.AddMaterial(mat)
	return mat, nil
}

func mergeIncludedVmtRecursive(base *keyvalues.KeyValue, includePath string) (*keyvalues.KeyValue, error) {
	parent, err := keyvalues2.ReadKeyValues(includePath)
	if err != nil {
		return base, errors.New("failed to read included vmt")
	}
	result, err := base.MergeInto(parent)
	if err != nil {
		return base, errors.New("failed to merge included vmt")
	}
	include, err := result.Find("include")
	if err == nil {
		newIncludePath, _ := include.AsString()
		if newIncludePath == includePath {
			err = result.RemoveChild("include")
			return &result, err
		}
		return mergeIncludedVmtRecursive(&result, newIncludePath)
	}
	return &result, nil
}
