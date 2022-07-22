package tool

import (
	"github.com/atotto/clipboard"
	"log"
)

func Copy(in string) {
	err := clipboard.WriteAll(in)
	if err != nil {
		log.Printf("Copy err:%s", err.Error())
	}
}
