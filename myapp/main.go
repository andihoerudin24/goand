package main

import (
	"github.com/andihoerudin24/goand"
	"myapp/handlers"
)

type application struct {
	App      *goand.Goand
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
