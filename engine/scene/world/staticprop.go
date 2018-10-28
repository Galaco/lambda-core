package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/bsp/primitives/game"
)

type StaticProp struct {
	entity.Base
	leafList []uint16
	model    *model.Model
}

func (prop *StaticProp) GetModel() *model.Model {
	return prop.model
}

func NewStaticProp(lumpProp game.IStaticPropDataLump, propLeafs *game.StaticPropLeafLump, renderable *model.Model) *StaticProp {
	prop := StaticProp{
		model: renderable,
	}
	for i := uint16(0); i < lumpProp.GetLeafCount(); i++ {
		prop.leafList = append(prop.leafList, propLeafs.Leaf[lumpProp.GetFirstLeaf()+i])
	}
	prop.Transform().Position = lumpProp.GetOrigin()
	prop.Transform().Rotation = lumpProp.GetAngles()

	return &prop
}
