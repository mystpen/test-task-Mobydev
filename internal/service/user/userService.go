package user

import (
	"github.com/mystpen/test-task-Mobydev/internal/model"
)

type UserRepo interface {
	GetUserEmail(string) error
	CreateUserDB(*model.CreateUserData) error
	CheckLoginDB(model.LoginUserData) (int, error)
	AddTokenDB(int, string) error
	GetUserByToken(string) (*model.User, error)
	GetUserInfoByID(int) (*model.UserInfo, error)
	ChangeUserInfo(*model.UserInfo) (*model.UserInfo, error)
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo}
}

func (u *UserService) CheckUserExists(user *model.CreateUserData) (bool, error) { //TODO: logs to service
	EmailExists := u.repo.GetUserEmail(user.Email)

	existBool := false

	if EmailExists == nil {
		existBool = true
	}
	if !existBool {
		err := u.CreateUser(user)
		if err != nil {
			return existBool, err
		}

	}
	return existBool, nil
}

func (u *UserService) CreateUser(user *model.CreateUserData) error {

	err := u.repo.CreateUserDB(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) CheckLogin(user model.LoginUserData) (int, error) {
	return u.repo.CheckLoginDB(user)
}

func (u *UserService) AddToken(userid int, cookie string) error {
	return u.repo.AddTokenDB(userid, cookie)
}

func (u *UserService) GetUserByToken(token string) (user *model.User, err error) {
	return u.repo.GetUserByToken(token)
}

func (u *UserService) GetUserInfo(user *model.User) (*model.UserInfo, error) {
	UserID := user.ID
	return u.repo.GetUserInfoByID(UserID)
}

func (u *UserService) ChangeUserInfo(user *model.UserInfo) (*model.UserInfo, error) {
	return u.repo.ChangeUserInfo(user)
}

func (u *UserService) GetUserInfoByID(UserID int64) (*model.UserInfo, error) {
	return u.repo.GetUserInfoByID(int(UserID))
}
