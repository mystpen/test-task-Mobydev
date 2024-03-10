package service

import (
	"github.com/mystpen/test-task-Mobydev/internal/repository"
	"github.com/mystpen/test-task-Mobydev/internal/service/user"
	"github.com/mystpen/test-task-Mobydev/internal/service/video"
)

type Service struct {
	UserService *user.UserService
	VideoService *video.VideoService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewUserService(repo.UserRepo),
		VideoService: video.NewVideoService(repo.VideoRepo),
	}
}
