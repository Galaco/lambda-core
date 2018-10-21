package material

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/filesystem"
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
	ResourceManager := filesystem.Manager()
	materialBasePath := "materials/"

	for _, materialPath := range materialList {
		vtfTexturePath := ""
		// Only load the filesystem once
		if ResourceManager.Get(materialBasePath+materialPath) == nil {
			if !readVmt(materialBasePath, materialPath) {
				debug.Log("Could not find: " + materialPath)
				missingList = append(missingList, materialPath)
				continue
			}
			vmt := ResourceManager.Get(materialPath).(*Vmt)

			// NOTE: in patch vmts include is not supported
			if vmt.GetProperty("baseTexture").AsString() != "" {
				vtfTexturePath = vmt.GetProperty("baseTexture").AsString() + ".vtf"
			}

			if vtfTexturePath != "" && !ResourceManager.Has(vtfTexturePath) {
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
	ResourceManager := filesystem.Manager()
	path := basePath + filePath + ".vmt"

	stream, err := filesystem.Load(path)
	if err != nil {
		return false
	}

	vmt, err := ParseVmt(filePath, stream)
	if err != nil {
		debug.Log(err)
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
		debug.Log(err)
		return false
	}
	// Store filesystem containing raw data in memory
	ResourceManager.Add(
		NewMaterial(
			filePath,
			texture,
			int(texture.GetHeader().Width),
			int(texture.GetHeader().Height)))

	// Finally generate the gpu buffer for the material
	ResourceManager.Get(filePath).(*Material).GenerateGPUBuffer()
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
