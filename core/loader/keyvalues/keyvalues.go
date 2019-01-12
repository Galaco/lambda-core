package keyvalues

import (
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/KeyValues"
)

// ReadKeyValues loads a keyvalues file.
// Its just a simple wrapper that combines the KeyValues library and
// the filesystem module.
func ReadKeyValues(filePath string) (*keyvalues.KeyValue, error) {
	stream, err := filesystem.GetFile(filePath)
	if err != nil {
		return nil, err
	}

	reader := keyvalues.NewReader(stream)
	kvs, err := reader.Read()

	return &kvs, err
}
