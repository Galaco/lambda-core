package material

import (
	"github.com/golang-source-engine/filesystem"
	"github.com/golang-source-engine/vmt"
)

func LoadVmtFromFilesystem(fs *filesystem.FileSystem, filePath string) (*vmt.Properties, error) {
	mat,err := vmt.FromFilesystem(filePath, fs, vmt.NewProperties())
	if err != nil {
		return nil, err
	}

	return mat.(*vmt.Properties), nil
}