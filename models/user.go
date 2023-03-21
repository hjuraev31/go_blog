package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
	Password  []byte `json: "-"`
	Phone     string `json: "phone"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}
