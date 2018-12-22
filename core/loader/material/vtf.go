package material

import (
	"github.com/galaco/Gource-Engine/core/filesystem"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/resource"
	"github.com/galaco/Gource-Engine/core/texture"
	"github.com/galaco/Gource-Engine/lib/vtf"
)

// LoadSingleTexture
func LoadSingleTexture(filePath string) texture.ITexture {
	if resource.Manager().GetTexture(materialRootPath+filePath) != nil {
		return resource.Manager().GetTexture(materialRootPath + filePath).(texture.ITexture)
	}
	if filePath == "" || !readVtf(materialRootPath+filePath) {
		return resource.Manager().GetTexture(resource.Manager().ErrorTextureName()).(texture.ITexture)
	}
	return resource.Manager().GetTexture(materialRootPath + filePath).(texture.ITexture)
}

func readVtf(path string) bool {
	ResourceManager := resource.Manager()
	stream, err := filesystem.GetFile(path)
	if err != nil {
		return false
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	read, err := vtf.ReadFromStream(stream)
	if err != nil {
		logger.Error(err)
		return false
	}
	// Store filesystem containing raw data in memory
	ResourceManager.AddTexture(
		texture.NewTexture2D(
			path,
			read,
			int(read.GetHeader().Width),
			int(read.GetHeader().Height)))

	// Finally generate the gpu buffer for the material
	ResourceManager.GetTexture(path).(texture.ITexture).Finish()
	return true
}
