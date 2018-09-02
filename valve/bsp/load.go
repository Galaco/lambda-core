package bsp

import (
	"github.com/galaco/bsp"
	"log"
)

func LoadBsp(filename string) *bsp.Bsp {
	file, err := bsp.ReadFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
