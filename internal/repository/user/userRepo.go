package user

import (
	"database/sql"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserStorage) CreateUserDB(user *model.CreateUserData) error {
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

func (u *UserStorage) AddTokenDB(userid int, cookieToken string) error {
	query := `UPDATE users
	SET token = ?, expires = DATETIME('now', '+6 hours')
	WHERE ? = id`
	if _, err := u.db.Exec(query, cookieToken, userid); err != nil {
		return err
	}
	return nil
}

func (u *UserStorage) RemoveTokenDB(token string) error {
	query := `UPDATE users
	SET token = NULL, expires = NULL
	WHERE token = ?`
	_, err := u.db.Exec(query, token)
	return err
}

func (u *UserStorage) CheckLoginDB(user model.LoginUserData) (int, error) {
	var hashedPassword string
	var userID int
	err := u.db.QueryRow("SELECT id, password FROM users WHERE username= $1", user.Email).Scan(
		&userID,
		&hashedPassword)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))

	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u *UserStorage) GetUserByToken(token string) (*model.User, error) {
	user := &model.User{}
	err := u.db.QueryRow("SELECT id, username, email FROM users WHERE token= $1", token).Scan(
		&user.ID,
		&user.Username,
		&user.Email)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByToken:")
	}
	return user, nil
}
