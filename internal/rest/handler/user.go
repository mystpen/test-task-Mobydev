package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mystpen/test-task-Mobydev/pkg"
)

type UserService interface{}

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	var inputUserData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//Decode json data
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
}
