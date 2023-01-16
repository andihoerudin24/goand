package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"myapp/data"
	"net/http"
	"strconv"
)

func (a *application) routes() *chi.Mux {
	//middleware must come before any routes

	//add routes here
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)

	a.App.Routes.Get("/users/login", a.Handlers.Userlogin)
	//a.App.Routes.Post("/users/login")

	a.App.Routes.Get("/create-user", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Andi",
			LastName:  "Hoerudin",
			Email:     "me@mail.com",
			Active:    1,
			Password:  "password",
		}
		id, err := a.Models.Users.Insert(u)
		if err != nil {
			a.App.ErrorLog.Println(err)
		}
		fmt.Fprintf(w, "%d: %s", id, u.FirstName)
	})

	a.App.Routes.Get("/get-all-users", func(writer http.ResponseWriter, request *http.Request) {
		users, err := a.Models.Users.GetAll()
		if err != nil {
			a.App.ErrorLog.Println(err)
		}
		for _, x := range users {
			fmt.Fprintf(writer, x.LastName)
		}
	})

	a.App.Routes.Get("/get-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, _ := a.Models.Users.Get(id)
		fmt.Fprintf(w, "%s %s %s", u.FirstName, u.LastName, u.Email)
	})

	a.App.Routes.Get("/update-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
		}

		u.LastName = a.App.RandomString(10)
		err = u.Update(*u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprintf(w, "update last name to %s", u.LastName)
	})

	//static route
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
