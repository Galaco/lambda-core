package gameinfo

import (
	"github.com/galaco/KeyValues"
	"io"
)

var gameInfo keyvalues.KeyValue

func Get() *keyvalues.KeyValue {
	return &gameInfo
}

func Load(stream io.Reader) (*keyvalues.KeyValue,error) {
	kvReader := keyvalues.NewReader(stream)

	kv, err := kvReader.Read()
	if err == nil {
		gameInfo = kv
	}

	return &gameInfo,err
}