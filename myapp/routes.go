package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) routes() *chi.Mux {
	//middleware must come before any routes

	//add routes here
	a.App.Routes.Get("/", a.Handlers.Home)

	a.App.Routes.Get("/jet", func(writer http.ResponseWriter, request *http.Request) {
		a.App.Render.JetPage(writer, request, "testjet", nil, nil)
	})

	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
