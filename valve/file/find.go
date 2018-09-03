package file

import (
	"bytes"
	"errors"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/vpk2"
	"io"
)

var gameVPK *vpk.VPK
var pakFile *lumps.Pakfile

func SetGameVPK(vpkfile *vpk.VPK) {
	gameVPK = vpkfile
}

func SetPakfile(pakfile *lumps.Pakfile) {
	pakFile = pakfile
}

func Load(filename string) (io.Reader,error) {
	// try to read from pakfile first
	f,err := pakFile.GetFile(filename)
	if err == nil && f != nil && len(f) != 0 {
		return bytes.NewReader(f),nil
	}

	// Fall back to game vpk
	entry := gameVPK.Entry(filename)
	if entry != nil {
		return entry.Open()
	}

	return nil, errors.New("Could not find: " + filename)
}
