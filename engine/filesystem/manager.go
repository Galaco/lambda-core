package filesystem

import "github.com/galaco/go-me-engine/engine/interfaces"

type manager struct {
	filelist map[string]interfaces.IFile
}

func (fm *manager) AddFile(file interfaces.IFile) {
	fm.filelist[file.GetFilePath()] = file
}

func (fm *manager) RemoveFile(filePath string) {
	delete(fm.filelist, filePath)
}

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