package scene

import (
	entity2 "github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/entity"
)

type Scene struct {
	world    *entity.WorldSpawn
	entities []entity2.IEntity
}

func (s *Scene) AddEntity(ent entity2.IEntity) {
	s.entities = append(s.entities, ent)
}

func (s *Scene) GetEntity(idx int) entity2.IEntity {
	if idx > len(s.entities) {
		return nil
	}
	return s.entities[idx]
}

func (s *Scene) NumEntities() int {
	return len(s.entities)
}

func (s *Scene) GetAllEntities() *[]entity2.IEntity {
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
