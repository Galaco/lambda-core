package mesh

import (
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/texture"
	"github.com/galaco/gosigl"
)

type Mesh struct {
	vertices            []float32
	normals             []float32
	textureCoordinates  []float32
	lightmapCoordinates []float32

	gpuInfo   buffer
	gpuObject *gosigl.VertexObject
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

func (mesh *Mesh) Finish() {
	if mesh.gpuInfo.IsPrepared == true {
		return
	}
	mesh.gpuObject = gosigl.NewMesh(mesh.vertices)
	gosigl.CreateVertexAttribute(mesh.gpuObject, mesh.textureCoordinates, 2)
	gosigl.CreateVertexAttribute(mesh.gpuObject, mesh.normals, 3)

	// @TODO Find a better solution
	if len(mesh.lightmapCoordinates) < 2 {
		mesh.lightmapCoordinates = []float32{0, 1}
	}
	gosigl.CreateVertexAttribute(mesh.gpuObject, mesh.lightmapCoordinates, 2)
	gosigl.FinishMesh()
	mesh.gpuInfo.FaceMode = gosigl.Triangles

	mesh.gpuInfo.IsPrepared = true
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

func (mesh *Mesh) Bind() {
	gosigl.BindMesh(mesh.gpuObject)
}

func (mesh *Mesh) Destroy() {
	gosigl.DeleteMesh(mesh.gpuObject)
}

func NewMesh() *Mesh {
	return &Mesh{}
}
