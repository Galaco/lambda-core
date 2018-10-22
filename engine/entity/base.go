package entity

import (
	"github.com/galaco/Gource-Engine/engine/core"
	entity2 "github.com/galaco/source-tools-common/entity"
)

// Base is an object in the game world.
// By itself entity is nothing more than an identifiable object located at a point in space
type Base struct {
	keyValues  *entity2.Entity
	transform  Transform

	handle     core.Handle
}

func (entity *Base) SetKeyValues(keyValues *entity2.Entity) {
	entity.keyValues = keyValues
}

func (entity *Base) KeyValues() *entity2.Entity {
	return entity.keyValues
}

func (entity *Base) Classname() string {
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

// Returns this entity's transform component
func (entity *Base) Transform() *Transform {
	return &entity.transform
}

func NewEntity(definition *entity2.Entity) Base {
	ent := Base{
		keyValues: definition,
		handle:    core.NewHandle(),
	}

	return ent
}
