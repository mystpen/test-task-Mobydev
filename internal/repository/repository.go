package repository

import (
	"database/sql"

	"github.com/mystpen/test-task-Mobydev/internal/repository/user"
	"github.com/mystpen/test-task-Mobydev/internal/repository/video"
)

type Repository struct {
	UserRepo *user.UserStorage
	VideoRepo *video.VideoStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: user.NewUserStorage(db),
		VideoRepo: video.NewVideoStorage(db),
	}
}
