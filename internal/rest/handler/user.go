package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/mystpen/test-task-Mobydev/pkg"
)

type UserService interface {
	CheckUserExists(*model.CreateUserData) (bool, error)
}

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
		h.Logger.ErrLog.Print(err.Error())
		return
	}
	if pkg.CheckEmail(inputUserData.Email) && pkg.CheckName(inputUserData.Username) && pkg.CheckPassword(inputUserData.Password) {

		user := &model.CreateUserData{
			Username: inputUserData.Username,
			Email:    inputUserData.Email,
			Password: inputUserData.Password,
		}

		existBool, err := h.UserService.CheckUserExists(user)
		if err != nil {
			pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			h.Logger.ErrLog.Print(err.Error())
			return
		}

		if !existBool {
			h.Logger.InfoLog.Print("user created")
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
		} else {
			message := "user already exists"
			h.Logger.ErrLog.Print(message)
			pkg.ErrorResponse(w, r, http.StatusConflict, err.Error())
		}
	} else {
		message := "incorrect format for email or username or password"
		h.Logger.ErrLog.Print(message)
		pkg.ErrorResponse(w, r, http.StatusBadRequest, message)
		return
	}

}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {

}
