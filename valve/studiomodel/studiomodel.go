package studiomodel

import (
	"github.com/galaco/go-me-engine/valve/studiomodel/mdl"
	"github.com/galaco/go-me-engine/valve/studiomodel/vtx"
	"github.com/galaco/go-me-engine/valve/studiomodel/vvd"
)

type StudioModel struct {
	Mdl *mdl.Mdl
	Vvd *vvd.Vvd
	Vtx *vtx.Vtx
}
