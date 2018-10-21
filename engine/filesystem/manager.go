package filesystem

import (
	"strings"
)

// Very generic filesystem storage.
// If the struct came from a filesystem, it should be obtainable from here
type manager struct {
	resources map[string]IFile
}

// Add a new filesystem
func (m *manager) Add(file IFile) {
	m.resources[strings.ToLower(file.GetFilePath())] = file
}

// Remove an open filesystem
func (m *manager) Remove(filePath string) {
	delete(m.resources, strings.ToLower(filePath))
}

// Find a specific filesystem
func (m *manager) Get(filePath string) IFile {
	return m.resources[strings.ToLower(filePath)]
}

func (m *manager) Has(filePath string) bool {
	return (m.resources[strings.ToLower(filePath)] != nil)
}

var resourceManager manager

func Manager() *manager {
	if resourceManager.resources == nil {
		resourceManager.resources = map[string]IFile{}
	}

	return &resourceManager
}
