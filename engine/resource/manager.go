package resource

import (
	"strings"
	"sync"
)

// Very generic filesystem storage.
// If the struct came from a filesystem, it should be obtainable from here
type manager struct {
	errorModelName   string
	errorTextureName string

	materials         map[string]IResource
	materialReadMutex sync.Mutex
	models            map[string]IResource
	modelReadMutex    sync.Mutex
}

// Add a new material
func (m *manager) AddMaterial(file IResource) {
	m.materialReadMutex.Lock()
	m.materials[strings.ToLower(file.GetFilePath())] = file
	m.materialReadMutex.Unlock()
}

// Add a new model
func (m *manager) AddModel(file IResource) {
	m.modelReadMutex.Lock()
	m.models[strings.ToLower(file.GetFilePath())] = file
	m.modelReadMutex.Unlock()
}

// Get Find a specific filesystem
func (m *manager) GetMaterial(filePath string) IResource {
	return m.materials[strings.ToLower(filePath)]
}

func (m *manager) GetModel(filePath string) IResource {
	return m.models[strings.ToLower(filePath)]
}

// ErrorModelName Get error model name
func (m *manager) ErrorModelName() string {
	return m.errorModelName
}

// SetErrorModelName Override the default error model.
// Useful for when HL2 assets are not available (they include the engine
// default model)
func (m *manager) SetErrorModelName(name string) {
	m.errorModelName = name
}

// ErrorTextureName Get error texture name
func (m *manager) ErrorTextureName() string {
	return m.errorTextureName
}

// SetErrorTextureName Override default error texture
func (m *manager) SetErrorTextureName(name string) {
	m.errorTextureName = name
}

// Has the specified file been loaded
func (m *manager) HasMaterial(filePath string) bool {
	m.materialReadMutex.Lock()
	if m.materials[strings.ToLower(filePath)] != nil {
		m.materialReadMutex.Unlock()
		return true
	}
	m.materialReadMutex.Unlock()
	return false
}

// Has the specified model been loaded
func (m *manager) HasModel(filePath string) bool {
	m.modelReadMutex.Lock()
	if m.models[strings.ToLower(filePath)] != nil {
		m.modelReadMutex.Unlock()
		return true
	}
	m.modelReadMutex.Unlock()
	return false
}

func (m *manager) Cleanup() {
	for _, mat := range m.materials {
		mat.Unload()
	}
	for _, model := range m.models {
		model.Unload()
	}
}

var resourceManager manager

// Manager returns the static filemanager
func Manager() *manager {
	if resourceManager.materials == nil {
		resourceManager.errorModelName = "models/error.mdl"
		resourceManager.errorTextureName = "materials/error.vtf"
		resourceManager.materials = map[string]IResource{}
		resourceManager.models = map[string]IResource{}
	}

	return &resourceManager
}
