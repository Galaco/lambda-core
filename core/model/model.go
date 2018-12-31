package model

import (
	"github.com/galaco/Gource-Engine/core/mesh"
)

// Model A collection of renderable primitives/submeshes
type Model struct {
	meshes   []mesh.IMesh
	fileName string
}

// AddMesh Add a new primitive
func (model *Model) AddMesh(meshes ...mesh.IMesh) {
	model.meshes = append(model.meshes, meshes...)
}

// GetMeshes Get all primitives/submeshes
func (model *Model) GetMeshes() []mesh.IMesh {
	return model.meshes
}

// Reset removes all meshes from this model
func (model *Model) Reset() {
	model.meshes = []mesh.IMesh{}
}

// GetFilePath returns where is model was found on disk
func (model *Model) GetFilePath() string {
	return model.fileName
}

//func (model *Model) Destroy() {
//	for _, m := range model.meshes {
//		m.Destroy()
//	}
//}

// NewModel returns a new Model
func NewModel(filename string, meshes ...mesh.IMesh) *Model {
	return &Model{
		fileName: filename,
		meshes:   meshes,
	}
}
