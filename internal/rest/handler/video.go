package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/mystpen/test-task-Mobydev/pkg"
)

type VideoService interface{
	GetVideoByID(int) (*model.VideoInfo, error)
	ChangeVideoInfo(*model.VideoInfo, int) (*model.VideoInfo, error)
}

func (h *Handler) getVideoInfo(w http.ResponseWriter, r *http.Request) {
	id, err := pkg.ReadIDParam(r)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusNotFound, err.Error())
		h.Logger.ErrLog.Print("Not Found")
		return
	}
	videoInfo, err := h.VideoService.GetVideoByID(int(id))
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}
	err = pkg.WriteJSON(w, http.StatusOK, pkg.Envelope{"video": videoInfo}, nil)
	if err != nil {
		h.Logger.ErrLog.Print(err)
		w.WriteHeader(500)
	}
}

func (h *Handler) putVideoInfo(w http.ResponseWriter, r *http.Request) {
	id, err := pkg.ReadIDParam(r)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusNotFound, err.Error())
		h.Logger.ErrLog.Print("Not Found")
		return
	}

	var inputVideoData struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Category    string `json:"category"`
		CreatedYear int    `json:"year"`
		Description string `json:"description"`
	}

	err = json.NewDecoder(r.Body).Decode(&inputVideoData)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}

	createdVideoInfo := &model.VideoInfo{
		Title:       inputVideoData.Title,
		Type:        inputVideoData.Type,
		Category:    inputVideoData.Category,
		CreatedYear: inputVideoData.CreatedYear,
		Description: inputVideoData.Description,
	}
	newVideoInfo, err := h.VideoService.ChangeVideoInfo(createdVideoInfo, int(id))
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}
	err = pkg.WriteJSON(w, http.StatusOK, pkg.Envelope{"video": newVideoInfo}, nil)
	if err != nil {
		h.Logger.ErrLog.Print(err)
		w.WriteHeader(500)
	}
}
