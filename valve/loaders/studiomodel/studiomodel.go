package studiomodel

import (
	"github.com/galaco/go-me-engine/valve/loaders/studiomodel/mdl"
	"github.com/galaco/go-me-engine/valve/loaders/studiomodel/vtx"
	"github.com/galaco/go-me-engine/valve/loaders/studiomodel/vvd"
)

type StudioModel struct {
	Mdl *mdl.Mdl
	Vvd *vvd.Vvd
	Vtx *vtx.Vtx
}
