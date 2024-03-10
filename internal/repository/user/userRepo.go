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

func (u *UserStorage) CreateUserDB(user *model.CreateUserData) error {
	_, err := u.db.Exec("INSERT INTO users (email, username, password, role) VALUES ($1, $2, $3, $4)",
		user.Email,
		user.Username,
		user.Password,
		"user")

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
	err := u.db.QueryRow("SELECT id, password FROM users WHERE email= $1", user.Email).Scan(
		&userID,
		&hashedPassword)
	if err != nil {
		return 0, err
	}
	if user.Password != hashedPassword {
		return 0, errors.New("incorrect password")
	}

	return userID, nil
}

func (u *UserStorage) GetUserByToken(token string) (*model.User, error) {
	user := &model.User{}
	err := u.db.QueryRow("SELECT id, username, email, role FROM users WHERE token= $1", token).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Role)
	if err != nil {
		return nil, errors.Wrap(err, "GetUserByToken:")
	}
	return user, nil
}

func (u *UserStorage) GetUserInfoByID(userID int) (*model.UserInfo, error) {
	userInfo := &model.UserInfo{}
	
	query := `SELECT id, username, email, COALESCE(phone_number, '')
	FROM users WHERE id= $1
	`
	err := u.db.QueryRow(query, userID).Scan(
		&userInfo.ID,
		&userInfo.Username,
		&userInfo.Email,
		&userInfo.Phone,
		)
	if err != nil {
		return userInfo, errors.Wrap(err, "user DB selecting:")
	}
	return userInfo, nil
}

func (u *UserStorage) ChangeUserInfo(createdUserInfo *model.UserInfo) (*model.UserInfo, error) {
	newUserInfo := &model.UserInfo{}
	query := `UPDATE users
	SET username=$1,
		phone_number=$2
	WHERE id = $3
	`
	_, err := u.db.Exec(query, 
		createdUserInfo.Username,
		createdUserInfo.Phone,
		createdUserInfo.ID,
		)
	if err != nil {
		return nil, errors.Wrap(err, "sql users:")
	}

	return newUserInfo, nil
}
