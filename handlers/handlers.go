package handlers

import (
	"encoding/json"
	"fmt"
	"gormapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllUsers ... handler that returns all users as json
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAllUsers())
}

//GetUser ... handler that returns a user as json
func GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := models.GetUser(id)
	if user.ID == 0 {
		fmt.Fprintln(w, 404)
		return
	}
	json.NewEncoder(w).Encode(user)
}

//PostUser ... makes a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	models.CreateUser(user)
}
