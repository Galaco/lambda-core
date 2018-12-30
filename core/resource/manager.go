package resource

import (
	"github.com/galaco/Gource-Engine/core/event"
	"github.com/galaco/Gource-Engine/core/filesystem"
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/galaco/Gource-Engine/core/resource/message"
	"github.com/galaco/Gource-Engine/core/texture"
	"strings"
	"sync"
)

// Very generic filesystem storage.
// If the struct came from a filesystem, it should be obtainable from here
type manager struct {
	errorModelName   string
	errorTextureName string

	materials         map[string]material.IMaterial
	materialReadMutex sync.Mutex
	textures          map[string]texture.ITexture
	textureReadMutex  sync.Mutex
	models            map[string]*model.Model
	modelReadMutex    sync.Mutex
}

// Add a new material
func (m *manager) AddMaterial(file material.IMaterial) {
	if m.HasMaterial(file.GetFilePath()) {
		return
	}
	m.materialReadMutex.Lock()
	m.materials[strings.ToLower(file.GetFilePath())] = file
	m.materialReadMutex.Unlock()

	event.GetEventManager().Dispatch(message.LoadedMaterial(file))
}

// Add a new material
func (m *manager) AddTexture(file texture.ITexture) {
	if m.HasTexture(file.GetFilePath()) {
		return
	}
	m.textureReadMutex.Lock()
	m.textures[strings.ToLower(file.GetFilePath())] = file
	m.textureReadMutex.Unlock()
	event.GetEventManager().Dispatch(message.LoadedTexture(file))
}

// Add a new model
func (m *manager) AddModel(file *model.Model) {
	if m.HasModel(file.GetFilePath()) {
		return
	}
	m.modelReadMutex.Lock()
	m.models[strings.ToLower(file.GetFilePath())] = file
	m.modelReadMutex.Unlock()
	event.GetEventManager().Dispatch(message.LoadedModel(file))
}

// Get Find a specific filesystem
func (m *manager) GetMaterial(filePath string) material.IMaterial {
	return m.materials[strings.ToLower(filePath)]
}

func (m *manager) GetTexture(filePath string) texture.ITexture {
	return m.textures[strings.ToLower(filePath)]
}

func (m *manager) GetModel(filePath string) *model.Model {
	return m.models[strings.ToLower(filePath)]
}

func (m *manager) Materials() map[string]material.IMaterial {
	return m.materials
}

func (m *manager) Textures() map[string]texture.ITexture {
	return m.textures
}

func (m *manager) Models() map[string]*model.Model {
	return m.models
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

func (m *manager) HasTexture(filePath string) bool {
	m.textureReadMutex.Lock()
	if m.textures[strings.ToLower(filePath)] != nil {
		m.textureReadMutex.Unlock()
		return true
	}
	m.textureReadMutex.Unlock()
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

func (m *manager) Empty() {
	for idx,val := range m.materials {
		event.GetEventManager().Dispatch(message.UnloadedMaterial(val))
		delete(m.materials, idx)
	}
	for idx,val := range m.textures {
		event.GetEventManager().Dispatch(message.UnloadedTexture(val))
		delete(m.textures, idx)
	}
	for idx,val := range m.models {
		event.GetEventManager().Dispatch(message.UnloadedModel(val))
		delete(m.models, idx)
	}
}

var resourceManager manager

// Manager returns the static filemanager
func Manager() *manager {
	if resourceManager.materials == nil {
		resourceManager.errorModelName = "models/error.mdl"
		resourceManager.errorTextureName = filesystem.BasePathMaterial + "error" + filesystem.ExtensionVtf
		resourceManager.materials = make(map[string]material.IMaterial, 1024)
		resourceManager.models = make(map[string]*model.Model, 256)
		resourceManager.textures = make(map[string]texture.ITexture, 256)
	}

	return &resourceManager
}
