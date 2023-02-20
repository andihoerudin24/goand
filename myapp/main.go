package main

import (
	"github.com/andihoerudin24/goand"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
)

type application struct {
	App        *goand.Goand
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
