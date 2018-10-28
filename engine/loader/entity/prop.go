package entity

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/loader/prop"
	"github.com/galaco/Gource-Engine/engine/model"
	entity2 "github.com/galaco/Gource-Engine/entity"
	"strings"
)

func DoesEntityReferenceStudioModel(ent entity.IEntity) bool {
	return strings.HasSuffix(ent.KeyValues().ValueForKey("model"), ".mdl")
}

func AssignStudioModelToEntity(entity entity.IEntity) {
	modelName := entity.KeyValues().ValueForKey("model")
	if !filesystem.Manager().Has(modelName) {
		m, _ := prop.LoadProp(modelName)
		entity.(entity2.IProp).SetModel(m)
	} else {
		entity.(entity2.IProp).SetModel(filesystem.Manager().Get(modelName).(*model.Model))
	}
}
