package users

import (
	"errors"

	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) ValidateUsername() bool {
	return len(user.Username) >= 2
}

func (user *User) ValidateEmail() bool {
	_, err := mail.ParseAddress(user.Email)
	return err == nil
}

func (user *User) ValidatePasswordAndConfirmation(confirmation string) bool {
	return bool(user.Password != "" && confirmation != "" && user.Password == confirmation)
}

func (user *User) HashPassword() []byte {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	return hashed
}

func CreateUser(userInfo UserDTO) (*User, error) {
	user := &User{}
	user.Email = userInfo.Email
	user.Username = userInfo.Username
	user.Password = userInfo.Password

	emailIsValid := true
	usernameIsValid := true

	if userInfo.Email != "" {
		emailIsValid = user.ValidateEmail()
	}

	if userInfo.Username != "" {
		usernameIsValid = user.ValidateUsername()
	}

	passwordIsValid := user.ValidatePasswordAndConfirmation(userInfo.ConfirmPassword)

	if emailIsValid && usernameIsValid && passwordIsValid {
		// Hashing the password
		user.Password = string(user.HashPassword())
		return user, nil
	}

	if !emailIsValid {
		return nil, errors.New("email is invalid")
	}

	if !usernameIsValid {
		return nil, errors.New("username is invalid")
	}

	if !passwordIsValid {
		return nil, errors.New("password is invalid")
	}

	return nil, errors.New("error to create a new user")

}
