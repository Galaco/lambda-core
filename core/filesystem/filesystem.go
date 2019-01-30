package filesystem

import (
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/vpk2"
)

type Path string

type FileSystem struct {
	gameVPKs         map[Path]vpk.VPK
	localDirectories []string
	pakFile          *lumps.Pakfile
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		gameVPKs:         map[Path]vpk.VPK{},
		localDirectories: make([]string, 0),
		pakFile:          nil,
	}
}
