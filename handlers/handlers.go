package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"userapi/models"

	"github.com/gorilla/mux"
)

//GetAllUsers ... handler that returns all users as json
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		fmt.Println("Caught an error")
		fmt.Fprintln(w, http.StatusInternalServerError)
	} else {
		fmt.Println(users)
		json.NewEncoder(w).Encode(users)
	}
}

//GetUser ... handler that returns a user as json
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := models.GetUser(id)

	if err != nil {
		fmt.Fprintln(w, http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

//PostUser ... makes a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	models.CreateUser(user)
}
