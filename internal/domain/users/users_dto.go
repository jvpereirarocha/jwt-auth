package users

type UserDTO struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

type Token struct {
	Token string
}
