package models

import (
	"log"
)

//User ... models for user
type User struct {
	ID       int    `json:"id"`
	Realname string `json:"realname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

//GetUser ... gets a user by id
func GetUser(id int) (User, error) {
	var user User
	row := DB.QueryRow("select id, realname, username, email from users where id = $1", id)
	err := row.Scan(&user.ID, &user.Realname, &user.Username, &user.Email)
	return user, err
}

//GetAllUsers ... gets all users
func GetAllUsers() ([]User, error) {
	rows, err := DB.Query("select id, realname, username, email from users")
	defer rows.Close()
	var users []User
	for rows.Next() {
		var tempUser User
		err := rows.Scan(&tempUser.ID, &tempUser.Realname, &tempUser.Username, &tempUser.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, tempUser)
	}
	return users, err
}

//CreateUser ... adds a new user, gets an http request with json body
func CreateUser(user User) {
	DB.Exec("insert into users (realname, username, email) values ($1,$2,$3)", user.Realname, user.Username, user.Email)
}

//UpdateUser
