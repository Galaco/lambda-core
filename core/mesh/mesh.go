package mesh

import (
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/texture"
)

type Mesh struct {
	vertices            []float32
	normals             []float32
	textureCoordinates  []float32
	lightmapCoordinates []float32

	material  material.IMaterial
	lightmap  texture.ITexture
}

func (mesh *Mesh) AddVertex(vertex ...float32) {
	mesh.vertices = append(mesh.vertices, vertex...)
}

func (mesh *Mesh) AddNormal(normal ...float32) {
	mesh.normals = append(mesh.normals, normal...)
}

func (mesh *Mesh) AddTextureCoordinate(uv ...float32) {
	mesh.textureCoordinates = append(mesh.textureCoordinates, uv...)
}

func (mesh *Mesh) AddLightmapCoordinate(uv ...float32) {
	mesh.lightmapCoordinates = append(mesh.lightmapCoordinates, uv...)
}

func (mesh *Mesh) Vertices() []float32 {
	return mesh.vertices
}

func (mesh *Mesh) Normals() []float32 {
	return mesh.normals
}

func (mesh *Mesh) TextureCoordinates() []float32 {
	return mesh.textureCoordinates
}

func (mesh *Mesh) LightmapCoordinates() []float32 {
	return mesh.lightmapCoordinates
}

func (mesh *Mesh) GetMaterial() material.IMaterial {
	return mesh.material
}

func (mesh *Mesh) SetMaterial(mat material.IMaterial) {
	mesh.material = mat
}

func (mesh *Mesh) GetLightmap() texture.ITexture {
	return mesh.lightmap
}

func (mesh *Mesh) SetLightmap(mat texture.ITexture) {
	mesh.lightmap = mat
}

func NewMesh() *Mesh {
	return &Mesh{}
}
