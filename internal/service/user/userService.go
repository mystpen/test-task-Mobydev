package user

import (
	"github.com/mystpen/test-task-Mobydev/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	GetUserEmail(string) error
	CreateUserDB(*model.CreateUserData) error
	CheckLoginDB(model.LoginUserData) (int, error)
	AddTokenDB(int, string) error
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
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPW)

	err = u.repo.CreateUserDB(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) CheckLogin(user model.LoginUserData) (int, error){
	return u.repo.CheckLoginDB(user)
}

func (u *UserService) AddToken(userid int, cookie string) error{
	return u.repo.AddTokenDB(userid, cookie)
}
