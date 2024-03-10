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

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
