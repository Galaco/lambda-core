package material

import (
	"errors"
	"github.com/galaco/Gource-Engine/core/filesystem"
	keyvalues2 "github.com/galaco/Gource-Engine/core/loader/keyvalues"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/resource"
	"github.com/galaco/Gource-Engine/core/texture"
	"github.com/galaco/KeyValues"
	"strings"
	"sync"
)

// LoadMaterialList GetFile all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialListConcurrent(materialList []string) {
	var wg sync.WaitGroup
	wg.Add(len(materialList))

	for _, mat := range materialList {
		go func() {
			defer wg.Done()
			loadMaterials(mat)
		}()
	}

	wg.Wait()
}
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
	errorMat.Textures.Albedo = ResourceManager.GetTexture(name).(texture.ITexture)
	ResourceManager.AddMaterial(errorMat)
}

// loadMaterials "private" function that actually does the loading
func loadMaterials(materialList ...string) (missingList []string) {
	ResourceManager := resource.Manager()

	for _, materialPath := range materialList {
		vtfTexturePath := ""

		if !strings.HasSuffix(materialPath, filesystem.ExtensionVmt) {
			materialPath += filesystem.ExtensionVmt
		}
		if ResourceManager.HasMaterial(filesystem.BasePathMaterial+materialPath) == true {
			continue
		}

		mat, err := readVmt(filesystem.BasePathMaterial + materialPath)
		if err != nil {
			logger.Warn("Failed to load material: %s. Reason: %s", filesystem.BasePathMaterial+materialPath, err)
			missingList = append(missingList, materialPath)
			continue
		}
		vmt := mat.(*material.Material)

		if vmt.BaseTextureName == "" {
			vmt.Textures.Albedo = ResourceManager.GetTexture(ResourceManager.ErrorTextureName()).(texture.ITexture)
			missingList = append(missingList, materialPath)
			continue
		}

		// NOTE: in patch vmts include is not supported
		vtfTexturePath = vmt.BaseTextureName
		if !strings.HasSuffix(vtfTexturePath, filesystem.ExtensionVtf) {
			vtfTexturePath = vtfTexturePath + filesystem.ExtensionVtf
		}

		vmt.Textures.Albedo = LoadSingleTexture(vtfTexturePath)
		if vmt.Textures.Albedo == nil {
			vmt.Textures.Albedo = ResourceManager.GetTexture(ResourceManager.ErrorTextureName()).(texture.ITexture)
			missingList = append(missingList, materialPath)
			continue
		}

		if vmt.BumpMapName != "" {
			vmt.Textures.Normal = LoadSingleTexture(vmt.BumpMapName)
		}
	}
	return missingList
}

// LoadSingleMaterial loads a single material with known file path
func LoadSingleMaterial(filePath string) material.IMaterial {
	if resource.Manager().HasMaterial(filesystem.BasePathMaterial + filePath) {
		return resource.Manager().GetMaterial(filesystem.BasePathMaterial + filePath).(material.IMaterial)
	}

	result := loadMaterials(filePath)
	if len(result) == 0 {
		return resource.Manager().GetMaterial(filesystem.BasePathMaterial + filePath).(material.IMaterial)

	}
	return resource.Manager().GetMaterial(resource.Manager().ErrorTextureName()).(material.IMaterial)
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
	if include != nil && err == nil {
		includePath, _ := include.AsString()
		root, err = mergeIncludedVmtRecursive(root, includePath)
		if err != nil {
			return nil, err
		}
	}

	// @NOTE this will be replaced with a proper kv->material builder
	mat,err := materialFromKeyValues(root, path)
	if err != nil {
		return nil,err
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

func materialFromKeyValues(kv *keyvalues.KeyValue, path string) (*material.Material,error) {
	shaderName := kv.Key()

	// $basetexture
	baseTexture := findKeyValueAsString(kv, "$basetexture")

	// $bumpmap
	bumpMapTexture := findKeyValueAsString(kv, "$bumpmap")


	return &material.Material{
		FilePath:        path,
		ShaderName:      shaderName,
		BaseTextureName: baseTexture,
		BumpMapName:     bumpMapTexture,
	}, nil
}

func findKeyValueAsString(kv *keyvalues.KeyValue, keyName string) string {
	target, err := kv.Find(keyName)
	if err != nil {
		return ""
	}
	ret, err := target.AsString()
	if err != nil {
		return ""
	}

	return ret
}
