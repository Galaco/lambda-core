package base

import (
	"github.com/galaco/go-me-engine/engine/core"
)

// Entity is an object in the game world.
// By itself entity is nothing more than an identifiable object located at a point in space
type Entity struct {
	handle core.Handle
	components []core.Handle
	transform TransformComponent
}

// Set this entity unique id
func (entity *Entity) SetHandle(handle core.Handle) {
	entity.handle = handle
}

// Return this entitys unique id
func (entity *Entity) GetHandle() core.Handle {
	return entity.handle
}

// Get all handles for this entity
func (entity *Entity) GetComponents() []core.Handle {
	return entity.components
}

// Associate a new component handle with this entity
func (entity *Entity) AddComponent(handle core.Handle) {
	entity.components = append(entity.components, handle)
}

// Returns this entity's transform component
func (entity *Entity) GetTransformComponent() *TransformComponent {
	return &entity.transform
}


func NewEntity() Entity {
	ent := Entity{
		handle: core.NewHandle(),
	}
	ent.GetTransformComponent().owner = core.NewHandle()

	return ent
}
