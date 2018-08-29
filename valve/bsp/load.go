package bsp

import (
	"log"
	"github.com/galaco/bsp"
)

func LoadBsp(filename string) *bsp.Bsp{
	file,err := bsp.ReadFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
	//
	//f, err := os.Open(filename)
	//if err!= nil {
	//	log.Fatal(err)
	//}
	//reader := bsp.NewReader(f)
	//
	//
	////Load file
	//f.Close()
	//
	//return file
}
