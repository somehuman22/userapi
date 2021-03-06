package main

import (
	"log"
	"net/http"
	"userapi/handlers"
	"userapi/models"

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
