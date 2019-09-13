package material

import (
	"errors"
	"github.com/galaco/KeyValues"
	filesystem2 "github.com/galaco/lambda-core/filesystem"
	"github.com/galaco/lambda-core/lib/util"
	keyvalues2 "github.com/galaco/lambda-core/loader/keyvalues"
	"github.com/galaco/lambda-core/material"
	"github.com/galaco/lambda-core/resource"
	"github.com/galaco/lambda-core/texture"
	"github.com/golang-source-engine/filesystem"
	"strings"
)

// LoadMaterialList GetFile all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(fs *filesystem.FileSystem, materialList []string) {
	loadMaterials(fs, materialList...)
}

// LoadErrorMaterial ensures that the error material has been loaded
func LoadErrorMaterial() {
	ResourceManager := resource.Manager()
	name := ResourceManager.ErrorTextureName()

	if ResourceManager.Material(name) != nil {
		return
	}

	// Ensure that error texture is available
	ResourceManager.AddTexture(texture.NewError(name))
	errorMat := material.NewMaterial(name)
	errorMat.Textures.Albedo = ResourceManager.Texture(name).(texture.ITexture)
	ResourceManager.AddMaterial(errorMat)
}

// loadMaterials "private" function that actually does the loading
func loadMaterials(fs *filesystem.FileSystem, materialList ...string) (missingList []string) {
	ResourceManager := resource.Manager()

	for _, materialPath := range materialList {
		vtfTexturePath := ""

		if !strings.HasSuffix(materialPath, filesystem2.ExtensionVmt) {
			materialPath += filesystem2.ExtensionVmt
		}
		if ResourceManager.HasMaterial(filesystem2.BasePathMaterial + materialPath) {
			continue
		}

		mat, err := readVmt(filesystem2.BasePathMaterial+materialPath, fs)
		if err != nil {
			util.Logger().Warn("Failed to load material: %s. Reason: %s", filesystem2.BasePathMaterial+materialPath, err)
			missingList = append(missingList, materialPath)
			continue
		}
		vmt := mat.(*material.Material)

		if vmt.BaseTextureName == "" {
			vmt.Textures.Albedo = ResourceManager.Texture(ResourceManager.ErrorTextureName()).(texture.ITexture)
			missingList = append(missingList, materialPath)

			ResourceManager.AddMaterial(vmt)
			continue
		}

		// NOTE: in patch vmts include is not supported
		vtfTexturePath = vmt.BaseTextureName
		if !strings.HasSuffix(vtfTexturePath, filesystem2.ExtensionVtf) {
			vtfTexturePath = vtfTexturePath + filesystem2.ExtensionVtf
		}

		vmt.Textures.Albedo = LoadSingleTexture(vtfTexturePath, fs)
		if vmt.Textures.Albedo == nil {
			vmt.Textures.Albedo = ResourceManager.Texture(ResourceManager.ErrorTextureName()).(texture.ITexture)
			missingList = append(missingList, materialPath)
			ResourceManager.AddMaterial(vmt)
			continue
		}

		if vmt.BumpMapName != "" {
			vmt.Textures.Normal = LoadSingleTexture(vmt.BumpMapName, fs)
		}
		ResourceManager.AddMaterial(vmt)
	}
	return missingList
}

// LoadSingleMaterial loads a single material with known file path
func LoadSingleMaterial(filePath string, fs *filesystem.FileSystem) material.IMaterial {
	if resource.Manager().HasMaterial(filesystem2.BasePathMaterial + filePath) {
		return resource.Manager().Material(filesystem2.BasePathMaterial + filePath).(material.IMaterial)
	}

	result := loadMaterials(fs, filePath)
	if len(result) == 0 {
		return resource.Manager().Material(filesystem2.BasePathMaterial + filePath).(material.IMaterial)

	}
	return resource.Manager().Material(resource.Manager().ErrorTextureName()).(material.IMaterial)
}

func readVmt(path string, fs *filesystem.FileSystem) (material.IMaterial, error) {
	kvs, err := keyvalues2.ReadKeyValues(path, fs)
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
		root, err = mergeIncludedVmtRecursive(root, includePath, fs)
		if err != nil {
			return nil, err
		}
	}

	// @NOTE this will be replaced with a proper kv->material builder
	mat, err := materialFromKeyValues(root, path)
	if err != nil {
		return nil, err
	}
	return mat, nil
}

func mergeIncludedVmtRecursive(base *keyvalues.KeyValue, includePath string, fs *filesystem.FileSystem) (*keyvalues.KeyValue, error) {
	parent, err := keyvalues2.ReadKeyValues(includePath, fs)
	if err != nil {
		return base, errors.New("failed to read included vmt")
	}
	parents, _ := parent.Children()
	result, err := base.Patch(parents[0])
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
		return mergeIncludedVmtRecursive(&result, newIncludePath, fs)
	}
	return &result, nil
}

func materialFromKeyValues(kv *keyvalues.KeyValue, path string) (*material.Material, error) {
	shaderName := kv.Key()

	// $basetexture
	baseTexture := findKeyValueAsString(kv, "$basetexture")

	// $bumpmap
	bumpMapTexture := findKeyValueAsString(kv, "$bumpmap")

	mat := material.NewMaterial(path)
	mat.ShaderName = shaderName
	mat.BaseTextureName = baseTexture
	mat.BumpMapName = bumpMapTexture
	return mat, nil
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
