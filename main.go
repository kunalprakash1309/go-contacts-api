package main

import (
	"github.com/go-contacts-api/app"
	"github.com/go-contacts-api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)

	srv := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: router,
	}
	log.Fatalln(srv.ListenAndServe())
}