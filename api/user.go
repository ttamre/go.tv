package api

import (
	"encoding/json"
	"net/http"
)


type User struct {
	ID       int 		`json:_id`		// primary key
	Username string		`json:username`
	Password string		`json:password`
	Reviews  []Review	`json:reviews`
}



// func CreateUser(username string, password string) (User, error) {
func CreateUser(r *http.Response, user *User) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// *user = User{
	// 	Username: r.FormValue("username"),
	// 	Password: r.FormValue("password")}
	// salt and hash password
	// add to database (INSERT INTO users (username, password) VALUES (username, password))
}

// func GetUser(username string, password string) (User, error) {
func GetUser(r *http.Response, user *User) {
	// get user from database (SELECT * FROM users WHERE username = username AND password = password)
	// if found
		// return user as a struct
	// else
		// return nil, errors.New("Invalid username or password")
	user = User{Username: r.FormValue("username"), Password: r.FormValue("password")}
}

func UpdateUsername(user *User, newUsername string) {
	// update database (UPDATE users SET username = newUsername WHERE id = user.ID)
	*user.Username = newUsername
}

func UpdatePassword(user *User, newPassword string) {
	// hash and salt password
	// update database (UPDATE users SET password = newPassword WHERE id = user.ID)
	*user.Password = newPassword
}

func DeleteUser(user *User) {
	// remove from database (DELETE FROM users WHERE id = user.ID)
	*user = nil
}
