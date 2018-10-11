package file

import (
	"bytes"
	"errors"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/vpk2"
	"io"
	"io/ioutil"
	"os"
)

var gameVPKs []*vpk.VPK
var fileSystemDirectories []string
var pakFile *lumps.Pakfile

func AddVpk(vpkFile *vpk.VPK) {
	gameVPKs = append(gameVPKs, vpkFile)
}

func AddSearchDirectory(directory string) {
	fileSystemDirectories = append(fileSystemDirectories, directory)
}

func SetPakfile(pakfile *lumps.Pakfile) {
	pakFile = pakfile
}

func Load(filename string) (io.Reader, error) {
	// try to read from pakfile first
	f, err := pakFile.GetFile(filename)
	if err == nil && f != nil && len(f) != 0 {
		return bytes.NewReader(f), nil
	}

	// Fall back to game vpk
	for _,fs := range gameVPKs {
		entry := fs.Entry(filename)
		if entry != nil {
			return entry.Open()
		}
	}

	// Fallback to local filesystem
	for _,dir := range fileSystemDirectories {
		if _, err := os.Stat(dir + filename); os.IsNotExist(err) {
			continue
		}
		file,err := ioutil.ReadFile(dir + filename)
		if err != nil {
			return nil,err
		}
		return bytes.NewBuffer(file), nil
	}


	return nil, errors.New("Could not find: " + filename)
}
