package model

type CreateUserData struct {
	Username string
	Email    string
	Password string
}

type LoginUserData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserInfo struct{
	ID int
	Username string
	Email string
	Phone string
}