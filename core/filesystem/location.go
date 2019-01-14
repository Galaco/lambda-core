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

// RegisterVpk registers a vpk package as a valid
// asset directory
func (fs *FileSystem) RegisterVpk(vpkFile *vpk.VPK) {
	fs.gameVPKs = append(fs.gameVPKs, *vpkFile)
}

func (fs *FileSystem) UnregisterVpk(vpkFile *vpk.VPK) {
	for idx := range fs.gameVPKs {
		if &fs.gameVPKs[idx] == vpkFile {
			fs.gameVPKs = append(fs.gameVPKs[:idx], fs.gameVPKs[idx+1:]...)
		}
	}
}

// RegisterLocalDirectory register a filesystem path as a valid
// asset directory
func (fs *FileSystem) RegisterLocalDirectory(directory string) {
	fs.localDirectories = append(fs.localDirectories, directory)
}

func (fs *FileSystem) UnregisterLocalDirectory(directory string) {
	for idx, dir := range fs.localDirectories {
		if dir == directory {
			if len(fs.localDirectories) == 1 {
				fs.localDirectories = make([]string, 0)
				return
			}
			fs.localDirectories = append(fs.localDirectories[:idx], fs.localDirectories[idx+1:]...)
		}
	}
}

// RegisterPakfile Set a pakfile to be used as an asset directory.
// This would normally be called during each map load
func (fs *FileSystem) RegisterPakFile(pakfile *lumps.Pakfile) {
	fs.pakFile = pakfile
}

// UnregisterPakfile removes the current pakfile from
// available search locations
func (fs *FileSystem) UnregisterPakFile() {
	fs.pakFile = nil
}

// GetFile attempts to get stream for filename.
// Search order is Pak->FileSystem->VPK
func (fs *FileSystem) GetFile(filename string) (io.Reader, error) {
	// sanitise file
	searchPath := NormalisePath(strings.ToLower(filename))

	// try to read from pakfile first
	if fs.pakFile != nil {
		f, err := fs.pakFile.GetFile(searchPath)
		if err == nil && f != nil && len(f) != 0 {
			return bytes.NewReader(f), nil
		}
	}

	// Fallback to local filesystem
	for _, dir := range fs.localDirectories {
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
	for _, vfs := range fs.gameVPKs {
		entry := vfs.Entry(searchPath)
		if entry != nil {
			return entry.Open()
		}
	}

	return nil, errors.New("could not find file")
}
