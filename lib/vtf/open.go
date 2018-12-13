package vtf

import (
	"github.com/galaco/vtf"
	"io"
)

// Basic wrapper around vtf library.

// ReadFromPath loads VMT from raw filepath
func ReadFromPath(filepath string) (*vtf.Vtf, error) {
	return vtf.ReadFromFile(filepath)
}

// ReadFromStream loads a vtf from a stream
func ReadFromStream(stream io.Reader) (*vtf.Vtf, error) {
	return vtf.ReadFromStream(stream)
}
