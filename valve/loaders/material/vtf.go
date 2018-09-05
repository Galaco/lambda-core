package material

import (
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	"github.com/galaco/go-me-engine/valve/file"
	"github.com/galaco/go-me-engine/valve/libwrapper/vtf"
	vpk2 "github.com/galaco/vpk2"
	"log"
)

// Load all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(pakData *lumps.Pakfile, vpkHandle *vpk2.VPK, materialList []string) {
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
			path := materialBasePath + materialPath + ".vmt"

			stream, err := file.Load(path)
			if err != nil {
				missingList = append(missingList, path)
				continue
			}

			vmt, err := ParseVmt(materialPath, stream)
			if err != nil {
				log.Println(err)
				continue
			}
			// Add file
			FileManager.AddFile(vmt)

			// NOTE: in patch vmts include is not supported
			if vmt.GetProperty("baseTexture").AsString() != "" {
				vtfTexturePath = vmt.GetProperty("baseTexture").AsString() + ".vtf"
			}
		}

		if vtfTexturePath != "" && FileManager.GetFile(vtfTexturePath) == nil {
			stream, err := file.Load(materialBasePath+vtfTexturePath)
			if err != nil {
				missingList = append(missingList, vtfTexturePath)
				continue
			}

			// Attempt to parse the vtf into color data we can use,
			// if this fails (it shouldn't) we can treat it like it was missing
			texture, err := vtf.ReadFromStream(stream)
			if err != nil {
				log.Println(err)
				continue
			}
			// Store file containing raw data in memory
			FileManager.AddFile(
				material.NewMaterial(
					vtfTexturePath,
					texture,
					int(texture.GetHeader().Width),
					int(texture.GetHeader().Height)))

			// Finally generate the gpu buffer for the material
			FileManager.GetFile(vtfTexturePath).(*material.Material).GenerateGPUBuffer()
		}
	}

	return missingList
}
