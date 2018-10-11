package studiomodel

import (
	"github.com/galaco/Gource-Engine/valve/loaders/studiomodel/mdl"
	"github.com/galaco/Gource-Engine/valve/loaders/studiomodel/vtx"
	"github.com/galaco/Gource-Engine/valve/loaders/studiomodel/vvd"
)

type StudioModel struct {
	Mdl *mdl.Mdl
	Vvd *vvd.Vvd
	Vtx *vtx.Vtx
}
