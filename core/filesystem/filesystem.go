package filesystem

import (
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/vpk2"
)

type FileSystem struct {
	gameVPKs         []vpk.VPK
	localDirectories []string
	pakFile          *lumps.Pakfile
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		gameVPKs:         make([]vpk.VPK, 0),
		localDirectories: make([]string, 0),
		pakFile:          nil,
	}
}
