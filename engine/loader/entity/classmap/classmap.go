package loader

import (
	entity3 "github.com/galaco/Gource-Engine/engine/entity"
)

type entityClassMapper struct {
	entityMap map[string]entity3.IEntity
}

func (classMap *entityClassMapper) Find(classname string) entity3.IEntity {
	if classMap.entityMap[classname] != nil {
		return classMap.entityMap[classname].New()
	}
	return nil
}

var classMap entityClassMapper

func RegisterClass(entity entity3.IClassname) {
	if classMap.entityMap == nil {
		classMap.entityMap = map[string]entity3.IEntity{}
	}

	classMap.entityMap[entity.Classname()] = entity.(entity3.IEntity)
}

func New(classname string) entity3.IEntity {
	return classMap.Find(classname)
}
