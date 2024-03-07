package api


type User struct {
	ID       int 		`json:_id`		// primary key
	Username string		`json:username`
	Password string		`json:password`
	Reviews  []Review	`json:reviews`
}



func CreateUser(username string, password string) (User, error) {
	// add to database
	return User{Username: username, Password: password, Reviews: []Review{}}, nil
}

func GetUser(username string, password string) (User, error) {
	// validate hash and salted password
	// if valid
		// get user from database
		// return user as a struct
	// else
		// return nil, errors.New("Invalid username or password")
	return User{Username: username, Password: password}, nil
}

func UpdateUsername(user *User, newUsername string) {
	// update database
	user.Username = newUsername
}

func UpdatePassword(user *User, newPassword string) {
	// update database
	// hash and salt password
	user.Password = newPassword
}

func DeleteUser(user *User) {
	// remove from database
	user = nil
}
