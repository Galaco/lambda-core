package scene

import (
	"github.com/galaco/Gource-Engine/engine/interfaces"
	"github.com/galaco/Gource-Engine/entity"
)

type Scene struct {
	world *entity.WorldSpawn
	entities []interfaces.IEntity
}

func (s *Scene) AddEntity(ent interfaces.IEntity) {
	s.entities = append(s.entities, ent)
}

func (s *Scene) GetEntity(idx int) interfaces.IEntity {
	if idx > len(s.entities) {
		return nil
	}
	return s.entities[idx]
}

func (s *Scene) NumEntities() int {
	return len(s.entities)
}

func (s *Scene) GetAllEntities() *[]interfaces.IEntity {
	return &s.entities
}

func (s *Scene) SetWorld(world *entity.WorldSpawn) {
	s.world = world
}

func (s *Scene) GetWorld() *entity.WorldSpawn {
	return s.world
}

var current Scene

func Get() *Scene {
	return &current
}