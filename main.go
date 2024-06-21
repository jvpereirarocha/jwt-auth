package main

import (
	"fmt"

	"github.com/jvpereirarocha/jwt-auth/internal/domain/users"
)

func main() {
	userInfos := users.UserDTO{Email: "teste@gmail.com", Username: "teste", Password: "hello", ConfirmPassword: "hello"}
	new_user, err := users.CreateUser(userInfos)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("User created! %v", new_user.Email)
}
