package filesystem

import (
	"strings"
	"sync"
)

// Very generic filesystem storage.
// If the struct came from a filesystem, it should be obtainable from here
type manager struct {
	resources map[string]IFile
	errorModelName string
	errorTextureName string

	readMutex sync.Mutex
}

// Add a new filesystem
func (m *manager) Add(file IFile) {
	m.readMutex.Lock()
	m.resources[strings.ToLower(file.GetFilePath())] = file
	m.readMutex.Unlock()
}

// Remove an open filesystem
func (m *manager) Remove(filePath string) {
	delete(m.resources, strings.ToLower(filePath))
}

// Get Find a specific filesystem
func (m *manager) Get(filePath string) IFile {
	return m.resources[strings.ToLower(filePath)]
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
func (m *manager) Has(filePath string) bool {
	m.readMutex.Lock()
	if m.resources[strings.ToLower(filePath)] != nil {
		m.readMutex.Unlock()
		return true
	}
	m.readMutex.Unlock()
	return false
}

var resourceManager manager

// Manager returns the static filemanager
func Manager() *manager {
	if resourceManager.resources == nil {
		resourceManager.errorModelName = "models/error.mdl"
		resourceManager.errorTextureName = "materials/error.vtf"
		resourceManager.resources = map[string]IFile{}
	}

	return &resourceManager
}
