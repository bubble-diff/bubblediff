package main

import (
	"log"

	"github.com/bubble-diff/bubblediff/config"
	"github.com/bubble-diff/bubblediff/db"
)

func main() {
	conf := config.Get()
	err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	r, err := InitRouter()
	if err != nil {
		log.Fatal(err)
	}
	err = r.Run(conf.ListenAddr)
	if err != nil {
		log.Fatal(err)
	}
}
