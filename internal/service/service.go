package service

import (
	"github.com/mystpen/test-task-Mobydev/internal/repository"
	"github.com/mystpen/test-task-Mobydev/internal/service/user"
)

type Service struct {
	UserService *user.UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
	}
}