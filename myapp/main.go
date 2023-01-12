package main

import (
	"github.com/andihoerudin24/goand"
	"myapp/data"
	"myapp/handlers"
)

type application struct {
	App      *goand.Goand
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
