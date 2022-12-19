package main

import (
	"github.com/andihoerudin24/goand"
	"log"
	"myapp/handlers"
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

	myHandlers := &handlers.Handlers{
		App: goands,
	}

	app := &application{
		App:      goands,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()
	return app
}
