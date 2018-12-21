package mesh

import (
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct {
	vertices            []float32
	normals             []float32
	textureCoordinates  []float32
	lightmapCoordinates []float32

	gpuInfo  buffer
	material material.IMaterial
	lightmap texture.ITexture
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
	// Gen vbo data
	gl.GenBuffers(1, &mesh.gpuInfo.Vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.Vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(mesh.vertices), gl.Ptr(mesh.vertices), gl.STATIC_DRAW)

	// gen vao data
	gl.GenVertexArrays(1, &mesh.gpuInfo.Vao)
	gl.BindVertexArray(mesh.gpuInfo.Vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.Vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// gen uv data
	gl.GenBuffers(1, &mesh.gpuInfo.UvBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.UvBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.textureCoordinates)*4, gl.Ptr(mesh.textureCoordinates), gl.STATIC_DRAW)

	// gen normal data
	gl.GenBuffers(1, &mesh.gpuInfo.NormalBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.NormalBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.normals)*4, gl.Ptr(mesh.normals), gl.STATIC_DRAW)

	// gen lightmap uv data
	gl.GenBuffers(1, &mesh.gpuInfo.LightmapUvBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.LightmapUvBuffer)
	// @TODO Find a better solution
	if len(mesh.lightmapCoordinates) < 2 {
		mesh.lightmapCoordinates = []float32{0, 1}
	}
	gl.BufferData(gl.ARRAY_BUFFER, len(mesh.lightmapCoordinates)*4, gl.Ptr(mesh.lightmapCoordinates), gl.STATIC_DRAW)

	mesh.gpuInfo.FaceMode = gl.TRIANGLES

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
	gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(mesh.gpuInfo.Vao)

	// UV's
	gl.EnableVertexAttribArray(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.UvBuffer)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	// Normals's
	gl.EnableVertexAttribArray(2)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.NormalBuffer)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)

	// Lightmap UV's
	gl.EnableVertexAttribArray(3)
	gl.BindBuffer(gl.ARRAY_BUFFER, mesh.gpuInfo.LightmapUvBuffer)
	gl.VertexAttribPointer(3, 2, gl.FLOAT, false, 0, nil)
}

func NewMesh() *Mesh {
	return &Mesh{}
}
