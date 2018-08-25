package vtf

import (
	"github.com/galaco/vtf"
	"io"
)

// Basic wrapper around vtf library.
func ReadFromPath(filepath string) (*vtf.Vtf, error){
	return vtf.ReadFromFile(filepath)
}
func ReadFromStream(stream io.Reader) (*vtf.Vtf, error){
	return vtf.ReadFromStream(stream)
}