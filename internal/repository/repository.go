package repository

import (
	"database/sql"

	"github.com/mystpen/test-task-Mobydev/internal/repository/user"
)

type Repository struct {
	UserRepo user.UserStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: *user.NewUserStorage(db),
	}
}
