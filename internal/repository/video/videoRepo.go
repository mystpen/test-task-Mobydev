package video

import "database/sql"

type VideoStorage struct {
	db *sql.DB
}

func NewVideoStorage(db *sql.DB) *VideoStorage {
	return &VideoStorage{db: db}
}
