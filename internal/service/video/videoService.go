package video

import "github.com/mystpen/test-task-Mobydev/internal/model"

type VideoRepo interface {
	GetVideoByID(int) (*model.VideoInfo, error)
	ChangeVideoInfo(*model.VideoInfo, int) (*model.VideoInfo, error)
}

type VideoService struct {
	repo VideoRepo
}

func NewVideoService(repo VideoRepo) *VideoService {
	return &VideoService{repo}
}

func (v *VideoService) GetVideoByID(videoID int) (*model.VideoInfo, error) {
	return v.repo.GetVideoByID(videoID)
}

func (v *VideoService) ChangeVideoInfo(createdVideoInfo *model.VideoInfo, videoID int) (*model.VideoInfo, error) {
	return v.repo.ChangeVideoInfo(createdVideoInfo, videoID)
}
