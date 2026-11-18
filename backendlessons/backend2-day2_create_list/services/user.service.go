package services

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
}

var user []Users
var lastID = 0

func CreateUser(username string, email string, phone int) Users {
	lastID++

	newuser := Users{
		Id:       lastID,
		Username: username,
		Email:    email,
		Phone:    phone,
	}

	user = append(user, newuser)

	return newuser
}

// List users
func List() []Users {
	return user
}
