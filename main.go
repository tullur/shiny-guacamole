package main

import (
	"aproj/controllers"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	pageC := controllers.NewStatic()
	usersController := controllers.NewUsers()

	r := mux.NewRouter()

	r.Handle("/", pageC.Home).Methods("GET")
	r.Handle("/contact", pageC.Contact).Methods("GET")
	r.Handle("/about", pageC.About).Methods("GET")

	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", r)
}
