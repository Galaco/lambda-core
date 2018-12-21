package cache

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
)

type InstancedProp struct {
	Name       string
	Skin       int
	Model      *model.Model
	Transforms []*entity.Transform

	transformationListCache []float32
	rebuild                 bool
}

func (instance *InstancedProp) TransformationMatrixList() *[]float32 {
	if instance.rebuild == false {
		return &instance.transformationListCache
	}
	//@TODO Recache
	return &instance.transformationListCache
}

func (instance *InstancedProp) Add(transform *entity.Transform) {
	instance.Transforms = append(instance.Transforms, transform)
	instance.rebuild = true
}

func NewInstancedProp(name string, skin int, prop *model.Model, transforms ...*entity.Transform) *InstancedProp {
	return &InstancedProp{
		Name:       name,
		Skin:       skin,
		Model:      prop,
		Transforms: transforms,
		rebuild:    true,
	}
}
