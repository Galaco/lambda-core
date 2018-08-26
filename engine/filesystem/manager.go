package filesystem

import "github.com/galaco/go-me-engine/engine/interfaces"

// Very generic file storage.
// If the struct came from a file, it should be obtainable from here
type manager struct {
	filelist map[string]interfaces.IFile
}

// Add a new file
func (fm *manager) AddFile(file interfaces.IFile) {
	fm.filelist[file.GetFilePath()] = file
}

// Remove an open file
func (fm *manager) RemoveFile(filePath string) {
	delete(fm.filelist, filePath)
}

// Find a specific file
func (fm *manager) GetFile(filePath string) interfaces.IFile {
	return fm.filelist[filePath]
}




var filemanager manager

func GetFileManager() *manager {
	if filemanager.filelist == nil {
		filemanager.filelist = map[string]interfaces.IFile{}
	}

	return &filemanager
}