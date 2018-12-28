package vtf

import (
	"github.com/galaco/vtf"
	"io"
)

// Basic wrapper around vtf library.

// ReadFromStream loads a vtf from a stream
func ReadFromStream(stream io.Reader) (*vtf.Vtf, error) {
	return vtf.ReadFromStream(stream)
}
