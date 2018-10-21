package gameinfo

import (
	"github.com/galaco/KeyValues"
	"os"
)

func LoadConfig(gameDirectory string) (*keyvalues.KeyValue, error) {
	// Load gameinfo.txt
	gameInfoFile, err := os.Open(gameDirectory + "/gameinfo.txt")
	if err != nil {
		return nil, err
	}
	return Load(gameInfoFile)
}
