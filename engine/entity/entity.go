package entity

import (
	"github.com/galaco/Gource-Engine/engine/component"
	"github.com/galaco/Gource-Engine/engine/core"
	entity2 "github.com/galaco/source-tools-common/entity"
)

// Base is an object in the game world.
// By itself entity is nothing more than an identifiable object located at a point in space
type Base struct {
	keyValues  *entity2.Entity
	handle     core.Handle
	components []core.Handle
	transform  component.TransformComponent
}

func (entity *Base) SetKeyValues(keyValues *entity2.Entity) {
	entity.keyValues = keyValues
}

func (entity *Base) KeyValues() *entity2.Entity {
	return entity.keyValues
}

func (entity *Base) ClassName() string {
	if entity.keyValues == nil {
		return ""
	}
	return entity.keyValues.ValueForKey("classname")
}

// Set this entity unique id
func (entity *Base) SetHandle(handle core.Handle) {
	entity.handle = handle
}

// Return this entitys unique id
func (entity *Base) GetHandle() core.Handle {
	return entity.handle
}

// Get all handles for this entity
func (entity *Base) GetComponents() []core.Handle {
	return entity.components
}

// Associate a new component handle with this entity
func (entity *Base) AddComponent(handle core.Handle) {
	entity.components = append(entity.components, handle)
}

// Returns this entity's transform component
func (entity *Base) GetTransformComponent() *component.TransformComponent {
	return &entity.transform
}

func NewEntity(definition *entity2.Entity) Base {
	ent := Base{
		keyValues: definition,
		handle:    core.NewHandle(),
	}
	ent.GetTransformComponent().SetOwnerHandle(core.NewHandle())

	return ent
}
