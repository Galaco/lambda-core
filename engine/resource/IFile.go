package resource

// Generic file object. If it was loaded from a path, it should
// implement this.
type IFile interface {
	GetFilePath() string
}
