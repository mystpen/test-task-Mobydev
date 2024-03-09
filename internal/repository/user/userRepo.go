package user

import (
	"database/sql"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/pkg/errors"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) GetUserEmail(userEmail string) error {
	var userID int
	err := u.db.QueryRow("SELECT id FROM users WHERE email= $1", userEmail).Scan(
		&userID)
	if err != nil {
		return errors.Wrap(err, "DB selecting")
	}

	return nil
}

func (u *UserStorage) CreateUserDB(user *model.CreateUserData) error{
	_, err := u.db.Exec("INSERT INTO users (email, username, password) VALUES ($1, $2, $3)",
		user.Email,
		user.Username,
		user.Password)

	// fmt.Println(user.Email, user.Username, user.PasswordHash) TODO:
	if err != nil {
		return errors.Wrap(err, "repository user insert")
	}
	return nil
}
