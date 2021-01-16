package main

import (
	"github.com/gorilla/mux"
	"github.com/go-contacts-api/app"
	"github.com/go-contacts-api/controllers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	


	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	
	srv := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: router,
	}
	log.Fatalln(srv.ListenAndServe())
}