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

func (mat *Material) Bind() {
	mat.Textures.BaseTexture.Bind()
}

func (mat *Material) Width() int {
	return mat.Textures.BaseTexture.Width()
}

func (mat *Material) Height() int {
	return mat.Textures.BaseTexture.Height()
}

func (mat Material) Destroy() {

}

func (mat Material) GetFilePath() string {
	return mat.FilePath
}
