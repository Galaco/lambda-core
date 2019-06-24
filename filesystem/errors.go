package filesystem

import (
	"fmt"
)

type FileNotFoundError struct {
	fileName string
}

func (err FileNotFoundError) Error() string {
	return fmt.Sprintf("%s not found in filesystem", err.fileName)
}

func NewFileNotFoundError(filename string) *FileNotFoundError {
	return &FileNotFoundError{
		fileName: filename,
	}
}
