package user

import "github.com/mystpen/test-task-Mobydev/internal/model"

type UserRepo interface {
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo}
}

func (u UserService) CheckUserExists(user *model.CreateUserData) (bool, error) {
	return true, nil
}
