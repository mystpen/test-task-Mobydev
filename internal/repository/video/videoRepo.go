package video

import (
	"database/sql"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/pkg/errors"
)

type VideoStorage struct {
	db *sql.DB
}

func NewVideoStorage(db *sql.DB) *VideoStorage {
	return &VideoStorage{db: db}
}

func (v *VideoStorage) GetVideoByID(ID int) (*model.VideoInfo, error) {
	videoInfo := &model.VideoInfo{}
	query := `SELECT id, title, type_project, category, created_year, description 
	FROM videos WHERE id= $1
	`
	err := v.db.QueryRow(query, ID).Scan(
		&videoInfo.ID,
		&videoInfo.Title,
		&videoInfo.Type,
		&videoInfo.Category,
		&videoInfo.CreatedYear,
		&videoInfo.Description,
		)
	if err != nil {
		return videoInfo, errors.Wrap(err, "videos DB selecting:")
	}
	return videoInfo, nil
}

func (v *VideoStorage) ChangeVideoInfo(createdVideoInfo *model.VideoInfo, ID int) (*model.VideoInfo, error) {
	query := `UPDATE videos
	SET title=$1,
		type_project=$2,
		category=$3,
		created_year=$4,
		description=$5
	WHERE id = $6
	RETURNING id
	`
	err := v.db.QueryRow(query, 
			createdVideoInfo.Title,
			createdVideoInfo.Type,
			createdVideoInfo.Category,
			createdVideoInfo.CreatedYear,
			createdVideoInfo.Description,
			ID,
		).Scan(&createdVideoInfo.ID)
	if err != nil {
		return nil, errors.Wrap(err, "sql videos:")
	}

	return createdVideoInfo, nil
}
