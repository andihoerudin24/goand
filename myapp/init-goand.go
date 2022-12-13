package main

import (
	"github.com/andihoerudin24/goand"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// init goand
	goands := &goand.Goand{}
	err = goands.New(path)
	if err != nil {
		log.Fatal(err)
	}
	goands.AppName = "myapp"
	goands.Debug = true

	app := &application{
		App: goands,
	}
	return app
}
