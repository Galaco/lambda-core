package scene

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/scene/world"
)

type Scene struct {
	world    world.World
	entities []entity.IEntity
	sky      *material.Cubemap

	cameras []entity.Camera
	currentCamera *entity.Camera
}

func (s *Scene) AddEntity(ent entity.IEntity) {
	s.entities = append(s.entities, ent)
}

func (s *Scene) GetEntity(idx int) entity.IEntity {
	if idx > len(s.entities) {
		return nil
	}
	return s.entities[idx]
}

func (s *Scene) FindEntitiesByKey(key string, value string) []entity.IEntity {
	ret := make([]entity.IEntity, 0)
	for idx,ent := range s.entities {
		if ent.KeyValues().ValueForKey(key) == value {
			ret = append(ret, s.entities[idx])
		}
	}
	return ret
}

func (s *Scene) NumEntities() int {
	return len(s.entities)
}

func (s *Scene) GetAllEntities() *[]entity.IEntity {
	return &s.entities
}

func (s *Scene) SetWorld(world *world.World) {
	s.world = *world
}

func (s *Scene) GetWorld() *world.World {
	return &s.world
}

func (s *Scene) GetSky() *material.Cubemap {
	return s.sky
}

func (s *Scene) AddCamera(camera *entity.Camera) {
	if s.cameras == nil {
		s.cameras = make([]entity.Camera, 0)
	}
	s.cameras = append(s.cameras, *camera)

	if s.currentCamera == nil {
		s.currentCamera = &s.cameras[0]
	}
}

func (s *Scene) CurrentCamera() *entity.Camera {
	return s.currentCamera
}

var currentScene Scene

func Get() *Scene {
	return &currentScene
}
