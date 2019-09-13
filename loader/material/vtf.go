package material

import (
	filesystem2 "github.com/galaco/lambda-core/filesystem"
	"github.com/galaco/lambda-core/lib/util"
	"github.com/galaco/lambda-core/resource"
	"github.com/galaco/lambda-core/texture"
	"github.com/galaco/vtf"
	"github.com/golang-source-engine/filesystem"
	"strings"
)

// LoadSingleTexture
func LoadSingleTexture(filePath string, fs *filesystem.FileSystem) texture.ITexture {
	filePath = filesystem.NormalisePath(filePath)
	if !strings.HasSuffix(filePath, filesystem2.ExtensionVtf) {
		filePath = filePath + filesystem2.ExtensionVtf
	}
	if resource.Manager().Texture(filesystem2.BasePathMaterial+filePath) != nil {
		return resource.Manager().Texture(filesystem2.BasePathMaterial + filePath).(texture.ITexture)
	}
	if filePath == "" {
		return resource.Manager().Texture(resource.Manager().ErrorTextureName()).(texture.ITexture)
	}
	mat, err := readVtf(filesystem2.BasePathMaterial+filePath, fs)
	if err != nil {
		util.Logger().Warn("Failed to load texture: %s. Reason: %s", filesystem2.BasePathMaterial+filePath, err)
		return resource.Manager().Texture(resource.Manager().ErrorTextureName()).(texture.ITexture)
	}
	return mat
}

// readVtf
func readVtf(path string, fs *filesystem.FileSystem) (texture.ITexture, error) {
	ResourceManager := resource.Manager()
	stream, err := fs.GetFile(path)
	if err != nil {
		return nil, err
	}

	// Attempt to parse the vtf into color data we can use,
	// if this fails (it shouldn't) we can treat it like it was missing
	read, err := vtf.ReadFromStream(stream)
	if err != nil {
		return nil, err
	}
	// Store filesystem containing raw data in memory
	ResourceManager.AddTexture(
		texture.NewTexture2D(
			path,
			read,
			int(read.Header().Width),
			int(read.Header().Height)))

	// Finally generate the gpu buffer for the material
	return ResourceManager.Texture(path).(texture.ITexture), nil
}
