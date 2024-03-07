package handler

import (
	"github.com/mystpen/test-task-Mobydev/internal/service"
)

type Handler struct {
	UserService UserService
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		UserService: service.UserService,
		// logger: logger.GetLoggerInstance(),
	}
}