package cache

import (
	"github.com/galaco/Gource-Engine/core/entity"
	"github.com/galaco/Gource-Engine/core/model"
)

type StaticProps struct {
	PropSkinCache []InstancedProp
}

func (list *StaticProps) All() *[]InstancedProp {
	return &list.PropSkinCache
}

func (list *StaticProps) Find(name string, skin int) *InstancedProp {
	for idx, c := range list.PropSkinCache {
		if c.Name == name && c.Skin == skin {
			return &list.PropSkinCache[idx]
		}
	}
	return nil
}

func (list *StaticProps) Add(name string, skin int, model *model.Model, transform *entity.Transform) *InstancedProp {
	c := list.Find(name, skin)
	if c != nil {
		c.Add(transform)
		return c
	}

	c = NewInstancedProp(name, skin, model, transform)
	list.PropSkinCache = append(list.PropSkinCache, *c)

	return list.Find(name, skin)
}
