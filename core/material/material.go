package material

import "github.com/galaco/Gource-Engine/core/texture"

type Material struct {
	ShaderName string
	Textures   struct {
		Albedo texture.ITexture
		Normal texture.ITexture
	}
	FilePath        string
	BaseTextureName string
	BumpMapName		string
	Properties      struct {
	}
}

// Width returns this materials width. Albedo is used to
// determine material width where possible
func (mat *Material) Width() int {
	return mat.Textures.Albedo.Width()
}

// Height returns this materials height. Albedo is used to
// determine material height where possible
func (mat *Material) Height() int {
	return mat.Textures.Albedo.Height()
}

// GetFilePath returns this materials location in whatever
// filesystem it was found
func (mat *Material) GetFilePath() string {
	return mat.FilePath
}
