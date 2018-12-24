package filesystem

import "strings"

func NormalisePath(filePath string) string {
	return strings.Replace(filePath, "\\", "/", -1)
}
