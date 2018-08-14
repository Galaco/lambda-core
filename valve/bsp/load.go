package bsp

import (
	"os"
	"log"
	"github.com/galaco/bsp"
)

func LoadBsp(filename string) *bsp.Bsp{
	f, err := os.Open(filename)
	if err!= nil {
		log.Fatal(err)
	}
	reader := bsp.NewReader(f)
	file,err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	//Load file
	f.Close()

	return file
}
