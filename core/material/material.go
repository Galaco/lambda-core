package material

import "github.com/galaco/Gource-Engine/core/texture"

type Material struct {
	ShaderName string
	Textures   struct {
		BaseTexture texture.ITexture
	}
	FilePath        string
	BaseTextureName string
	Properties      struct {
	}
}

// Bind is used to ready all textures this material references
func (mat *Material) Bind() {
	mat.Textures.BaseTexture.Bind()
}

// Width returns this materials width. BaseTexture is used to
// determine material width where possible
func (mat *Material) Width() int {
	return mat.Textures.BaseTexture.Width()
}

// Height returns this materials height. BaseTexture is used to
// determine material height where possible
func (mat *Material) Height() int {
	return mat.Textures.BaseTexture.Height()
}

// Destroy does any cleanup required to correctly deallocate this material
func (mat Material) Destroy() {

}

// GetFilePath returns this materials location in whatever
// filesystem it was found
func (mat Material) GetFilePath() string {
	return mat.FilePath
}
