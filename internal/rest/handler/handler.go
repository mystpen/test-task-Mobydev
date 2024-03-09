package handler

import (
	"github.com/mystpen/test-task-Mobydev/internal/logger"
	"github.com/mystpen/test-task-Mobydev/internal/service"
)

type Handler struct {
	UserService UserService
	Logger      *logger.Logger
}

func NewHandler(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		UserService: service.UserService,
		Logger:      logger,
	}
}
