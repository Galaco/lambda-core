package material

import (
	"log"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	vpk2 "github.com/galaco/vpk2"
	"github.com/galaco/go-me-engine/valve/libwrapper/vtf"
	"path/filepath"
	"github.com/galaco/bsp/lumps"
	"bytes"
)

// Load all materials referenced in the map
// NOTE: There is a priority:
// 1. BSP pakfile
// 2. Game directory
// 3. Game VPK
// 4. Other game shared VPK
func LoadMaterialList(pakData *lumps.Pakfile, vpkHandle *vpk2.VPK, materialList []string) {
	missing := readFromPakfile(pakData, materialList)

	missing = readFromVPK(vpkHandle, missing)

	for _,path := range missing {
		log.Println("Could not find: " + path)
	}
}

func readFromPakfile(pakData *lumps.Pakfile, materialList []string) (missingList []string) {
	FileManager := filesystem.GetFileManager()
	materialBasePath := "materials/"

	for _,materialPath := range materialList {
		vtfTexturePath := ""
		// Only load the file once
		if FileManager.GetFile(materialBasePath + materialPath) == nil {
			// Pakfiles store full filepath inc. extension
			pakFileDir := materialBasePath + materialPath + ".vmt"
			vmtData, err := pakData.GetFile(pakFileDir)
			if err != nil || len(vmtData) == 0 {
				// If not found, just let it bubble up to other file finding mechanisms
				missingList = append(missingList, materialPath)
				continue
			}
			// Import vmt
			vmtStream := bytes.NewReader(vmtData)
			vmt,err := ParseVmt(materialPath, vmtStream)
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

		// Load $baseTexture
		if vtfTexturePath != "" && FileManager.GetFile(vtfTexturePath) == nil {
			vtfData, err := pakData.GetFile(materialBasePath + vtfTexturePath)
			if err != nil || len(vtfData) == 0 {
				missingList = append(missingList, materialPath + ".vtf")
				continue
			}
			vtfStream := bytes.NewReader(vtfData)

			// Attempt to parse the vtf into color data we can use,
			// if this fails (it shouldn't) we can treat it like it was missing
			texture,err := vtf.ReadFromStream(vtfStream)
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

func readFromVPK(vpkHandle *vpk2.VPK, materialList []string) (missingList []string) {
	FileManager := filesystem.GetFileManager()
	materialBasePath := "materials/"


	for _,materialPath := range materialList {
		var vmt *Vmt
		vtfTexturePath := ""
		// If the material is a vtf, skip over vmt reading,
		if filepath.Ext(materialPath) != ".vtf" {
			// Ensure we've loaded the .vmt file
			vmtPath := materialPath + ".vmt"
			if FileManager.GetFile(vmtPath) == nil {
				vmtF := vpkHandle.Entry(materialBasePath + vmtPath)
				if vmtF != nil {
					f2,err := vmtF.Open()
					vmt,err = ParseVmt(materialPath, f2)
					if err != nil {
						log.Println(err)
						continue
					}
				} else {
					missingList = append(missingList, materialPath)
					continue
				}

				FileManager.AddFile(vmt)
				vtfTexturePath =  vmt.GetProperty("baseTexture").AsString() + ".vtf"
			}
		} else {
			vtfTexturePath = materialPath
		}

		// Load $baseTexture
		if FileManager.GetFile(vtfTexturePath) == nil {
			// Load file from vpk into memory
			vpkFile := vpkHandle.Entry(materialBasePath + vtfTexturePath)
			if vpkFile == nil {
				log.Println("Couldn't find vtf: " + vtfTexturePath)
				missingList = append(missingList, vtfTexturePath)
				continue
			}
			file,err := vpkFile.Open()

			// Its quite possible for a texture to be missing, just skip it.
			if err != nil {
				log.Println("Couldn't open vtf: " + vtfTexturePath)
				continue
			}

			// Attempt to parse the vtf into color data we can use,
			// if this fails (it shouldn't) we can treat it like it was missing
			texture,err := vtf.ReadFromStream(file)
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
