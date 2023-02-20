package main

import (
	"github.com/andihoerudin24/goand"
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
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

	myMiddleware := &middleware.Middleware{
		App: goands,
	}

	myHandlers := &handlers.Handlers{
		App: goands,
	}

	app := &application{
		App:        goands,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models
	return app
}
