package filesystem

import (
	"bytes"
	"errors"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/vpk2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var gameVPKs []*vpk.VPK
var localDirectories []string
var pakFile *lumps.Pakfile

// RegisterVpk registers a vpk package as a valid
// asset directory
func RegisterVpk(vpkFile *vpk.VPK) {
	gameVPKs = append(gameVPKs, vpkFile)
}

func UnregisterVpk(vpkFile *vpk.VPK) {
	for idx, pkg := range gameVPKs {
		if pkg == vpkFile {
			gameVPKs = append(gameVPKs[:idx], gameVPKs[idx+1:]...)
		}
	}
}

// RegisterLocalDirectory register a filesystem path as a valid
// asset directory
func RegisterLocalDirectory(directory string) {
	localDirectories = append(localDirectories, directory)
}

func UnregisterLocalDirectory(directory string) {
	for idx, dir := range localDirectories {
		if dir == directory {
			localDirectories = append(localDirectories[:idx], localDirectories[idx+1:]...)
		}
	}
}

// RegisterPakfile Set a pakfile to be used as an asset directory.
// This would normally be called during each map load
func RegisterPakfile(pakfile *lumps.Pakfile) {
	pakFile = pakfile
}

func UnregisterPakfile() {
	pakFile = nil
}

// GetFile attempts to get stream for filename.
// Search order is Pak->FileSystem->VPK
func GetFile(filename string) (io.Reader, error) {
	// sanitise file
	searchPath := strings.ToLower(filename)

	// try to read from pakfile first
	if pakFile != nil {
		f, err := pakFile.GetFile(searchPath)
		if err == nil && f != nil && len(f) != 0 {
			return bytes.NewReader(f), nil
		}
	}

	// Fallback to local filesystem
	for _, dir := range localDirectories {
		if _, err := os.Stat(dir + "\\" + searchPath); os.IsNotExist(err) {
			continue
		}
		file, err := ioutil.ReadFile(dir + searchPath)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(file), nil
	}

	// Fall back to game vpk
	for _, fs := range gameVPKs {
		entry := fs.Entry(searchPath)
		if entry != nil {
			return entry.Open()
		}
	}

	return nil, errors.New("Could not find: " + searchPath)
}
