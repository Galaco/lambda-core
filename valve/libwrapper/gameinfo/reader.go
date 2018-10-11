package gameinfo

import (
	"github.com/galaco/vmf"
	"io"
)

var gameInfo vmf.Vmf

func Get() *vmf.Vmf {
	return &gameInfo
}

func Load(stream io.Reader) (*vmf.Vmf,error) {
	kvReader := vmf.NewReader(stream)

	kv, err := kvReader.Read()
	if err == nil {
		gameInfo = kv
	}

	return &gameInfo,err
}