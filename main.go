package main

import (
	"aproj/controllers"
	"aproj/models"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "adb"
)

func main() {
	pgConnection := fmt.Sprintf("host=%s port=%d user=%s password = %s dbname = %s sslmode = disable",
		host, port, user, password, dbname)

	us, err := models.NewUserService(pgConnection)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer us.Close()
	us.AutoMigrate()

	pageC := controllers.NewStatic()
	usersController := controllers.NewUsers(us)

	r := mux.NewRouter()

	r.Handle("/", pageC.Home).Methods("GET")
	r.Handle("/contact", pageC.Contact).Methods("GET")
	r.Handle("/about", pageC.About).Methods("GET")

	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")

	r.Handle("/login", usersController.LogInView).Methods("GET")
	r.HandleFunc("/login", usersController.Login).Methods("POST")

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", r)
}
