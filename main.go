package main

import (
	"gormapi/handlers"
	"gormapi/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.InitDB()
	router := mux.NewRouter()
	models.InitDB()

	//assign handlers
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/new", handlers.PostUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}
