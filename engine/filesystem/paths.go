package filesystem

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

// AddVpk registers a vpk package as a valid
// asset directory
func AddVpk(vpkFile *vpk.VPK) {
	gameVPKs = append(gameVPKs, vpkFile)
}

// AddSearchDirectory register a filesystem path as a valid
// asset directory
func AddSearchDirectory(directory string) {
	fileSystemDirectories = append(fileSystemDirectories, directory)
}

// SetPakfile Set a pakfile to be used as an asset directory.
// This would normally be called during each map load
func SetPakfile(pakfile *lumps.Pakfile) {
	pakFile = pakfile
}

// Load attempts to get stream for filename.
// Search order is Pak->VPK->FileSystem
func Load(filename string) (io.Reader, error) {
	// try to read from pakfile first
	f, err := pakFile.GetFile(filename)
	if err == nil && f != nil && len(f) != 0 {
		return bytes.NewReader(f), nil
	}

	// Fall back to game vpk
	for _, fs := range gameVPKs {
		entry := fs.Entry(filename)
		if entry != nil {
			return entry.Open()
		}
	}

	// Fallback to local filesystem
	for _, dir := range fileSystemDirectories {
		if _, err := os.Stat(dir + filename); os.IsNotExist(err) {
			continue
		}
		file, err := ioutil.ReadFile(dir + filename)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(file), nil
	}

	return nil, errors.New("Could not find: " + filename)
}
