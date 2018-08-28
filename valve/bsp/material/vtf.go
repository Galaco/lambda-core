package material

import (
	"log"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	vpk2 "github.com/galaco/vpk2"
	"github.com/galaco/go-me-engine/valve/libwrapper/vtf"
)

func LoadMaterialList(vpkHandle *vpk2.VPK, materialList []string) {
	FileManager := filesystem.GetFileManager()
	materialBasePath := "materials/"

	var vmt *Vmt

	for _,materialPath := range materialList {
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
				log.Println("Couldn't open vmt: " + materialPath)
				continue
			}

			FileManager.AddFile(vmt)
		}

		// Load $baseTexture
		vtfTexturePath :=  vmt.GetProperty("baseTexture").AsString() + ".vtf"
		if FileManager.GetFile(vtfTexturePath) == nil {
			// Load file from vpk into memory
			vpkFile := vpkHandle.Entry(materialBasePath + vtfTexturePath)
			if vpkFile == nil {
				log.Println("Couldn't find vtf: " + vtfTexturePath)
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
		}

		// Finally generate the gpu buffer for the material
		FileManager.GetFile(vtfTexturePath).(*material.Material).GenerateGPUBuffer()
	}
}
