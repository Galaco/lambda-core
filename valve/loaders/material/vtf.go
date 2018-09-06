package material

import (
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	"github.com/galaco/go-me-engine/valve/file"
	"github.com/galaco/go-me-engine/valve/libwrapper/vtf"
	"log"
)

// Load all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(materialList []string) {
	missing := read(materialList)

	for _, path := range missing {
		log.Println("Could not find: " + path)
	}
}

func read(materialList []string) (missingList []string) {
	FileManager := filesystem.GetFileManager()
	materialBasePath := "materials/"

	for _, materialPath := range materialList {
		vtfTexturePath := ""
		// Only load the file once
		if FileManager.GetFile(materialBasePath+materialPath) == nil {
			if !readVmt(materialBasePath, materialPath) {
				missingList = append(missingList, materialPath)
				continue
			}
			vmt := FileManager.GetFile(materialPath).(*Vmt)

			// NOTE: in patch vmts include is not supported
			if vmt.GetProperty("baseTexture").AsString() != "" {
				vtfTexturePath = vmt.GetProperty("baseTexture").AsString() + ".vtf"
			}
		}

		if vtfTexturePath != "" && FileManager.GetFile(vtfTexturePath) == nil {
			if !readVtf(materialBasePath, vtfTexturePath) {
				missingList = append(missingList, vtfTexturePath)
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
		log.Println(err)
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
		log.Println(err)
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

	for _,ext := range exts {
		readVtf("materials/skybox/", skyName + ext + ".vtf")
	}
}