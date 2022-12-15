package main

import "github.com/andihoerudin24/goand"

type application struct {
	App *goand.Goand
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
