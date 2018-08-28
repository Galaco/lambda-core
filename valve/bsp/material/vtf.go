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

	for _,materialPath := range materialList {
		var vmt *Vmt
		vmtF := vpkHandle.Entry("materials/" + materialPath + ".vmt")
		if vmtF != nil {
			f2,err := vmtF.Open()
			vmt,err = ParseVmt(materialPath, f2)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println("Could not open: " + materialPath)
			continue
		}

		log.Println(vmt.GetProperty("baseTexture").AsString())

		if FileManager.GetFile(materialPath) == nil {
			// Load file from vpk into memory
			vpkFile := vpkHandle.Entry("materials/" + vmt.GetProperty("baseTexture").AsString() + ".vtf")
			if vpkFile == nil {
				log.Println("Couldnt find material: materials/" + materialPath + ".vtf")
				continue
			}
			file,err := vpkFile.Open()

			// Its quite possible for a texture to be missing, just skip it.
			if err != nil {
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
					materialPath,
					texture,
					int(texture.GetHeader().Width),
					int(texture.GetHeader().Height)))
		}

		// Finally generate the gpu buffer for the material
		FileManager.GetFile(materialPath).(*material.Material).GenerateGPUBuffer()
	}
}
