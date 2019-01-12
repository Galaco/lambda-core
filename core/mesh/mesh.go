package mesh

import (
	"github.com/galaco/Lambda-Core/core/material"
	"github.com/galaco/Lambda-Core/core/mesh/util"
	"github.com/galaco/Lambda-Core/core/texture"
)

type Mesh struct {
	vertices            []float32
	normals             []float32
	uvs                 []float32
	tangents            []float32
	lightmapCoordinates []float32

	material material.IMaterial
	lightmap texture.ITexture
}

func (mesh *Mesh) AddVertex(vertex ...float32) {
	mesh.vertices = append(mesh.vertices, vertex...)
}

func (mesh *Mesh) AddNormal(normal ...float32) {
	mesh.normals = append(mesh.normals, normal...)
}

func (mesh *Mesh) AddUV(uv ...float32) {
	mesh.uvs = append(mesh.uvs, uv...)
}

func (mesh *Mesh) AddTangent(tangent ...float32) {
	mesh.tangents = append(mesh.tangents, tangent...)
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

func (mesh *Mesh) UVs() []float32 {
	return mesh.uvs
}

func (mesh *Mesh) Tangents() []float32 {
	return mesh.tangents
}

func (mesh *Mesh) LightmapCoordinates() []float32 {
	// use standard uvs if there is no lightmap. Not ideal,
	// but there MUST be UVs, but they'll be ignored anyway if there is no
	// lightmap
	if len(mesh.lightmapCoordinates) == 0 {
		return mesh.UVs()
	}
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

func (mesh *Mesh) GenerateTangents() {
	mesh.tangents = util.GenerateTangents(mesh.Vertices(), mesh.Normals(), mesh.UVs())
}

func NewMesh() *Mesh {
	return &Mesh{}
}
