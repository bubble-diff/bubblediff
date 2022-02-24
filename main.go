package main

import (
	"log"

	"github.com/bubble-diff/bubblediff/app"
	"github.com/bubble-diff/bubblediff/config"
)

func main() {
	conf := config.Get()
	err := app.Init()
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
