package model

import (
	"github.com/galaco/Gource-Engine/core/entity"
	"github.com/galaco/bsp/primitives/game"
)

// StaticProp is a somewhat specialised model
// that implements a few core entity features (largely because
// it is basically a renderable entity that cannot do anything or be reference)
type StaticProp struct {
	entity.Base
	leafList []uint16
	model    *Model
}

// GetModel returns props model
func (prop *StaticProp) GetModel() *Model {
	return prop.model
}

// LeafList returrns all leafs that this props is in
func (prop *StaticProp) LeafList() []uint16 {
	return prop.leafList
}

// NewStaticProp returns new StaticProp
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
