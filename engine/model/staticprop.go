package model

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/bsp/primitives/game"
)

type StaticProp struct {
	entity.Base
	leafList []uint16
	model    *Model
}

func (prop *StaticProp) GetModel() *Model {
	return prop.model
}

func (prop *StaticProp) LeafList() []uint16 {
	return prop.leafList
}

func NewStaticProp(lumpProp game.IStaticPropDataLump, propLeafs *game.StaticPropLeafLump, renderable *Model) *StaticProp {
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
