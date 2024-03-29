package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/mystpen/test-task-Mobydev/pkg"
)

type UserService interface {
	CheckUserExists(*model.CreateUserData) (bool, error)
	CheckLogin(model.LoginUserData) (int, error)
	AddToken(int, string) error
	GetUserByToken(string) (*model.User, error)
	GetUserInfo(*model.User) (*model.UserInfo, error)
	ChangeUserInfo(*model.UserInfo) (*model.UserInfo, error)
	GetUserInfoByID(int64) (*model.UserInfo, error)
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
			env := pkg.Envelope{
				"status": "success",
				"info":   "user created",
			}
			pkg.WriteJSON(w, http.StatusCreated, env, nil)
		} else {
			message := "user already exists"
			h.Logger.ErrLog.Print(message)
			pkg.ErrorResponse(w, r, http.StatusConflict, message)
		}
	} else {
		message := "incorrect format for email or username or password"
		h.Logger.ErrLog.Print(message)
		pkg.ErrorResponse(w, r, http.StatusBadRequest, message)
		return
	}

}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	var inputUserData model.LoginUserData
	//Decode json data
	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}

	userid, err := h.UserService.CheckLogin(inputUserData)
	if err == nil {
		cookieToken := pkg.SetCookie(w)
		err = h.UserService.AddToken(userid, cookieToken)
		if err != nil {
			h.Logger.ErrLog.Print(err.Error())
			return
		}

		h.Logger.InfoLog.Print("token added")
	} else {
		pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}
}

func (h *Handler) getUserInfo(w http.ResponseWriter, r *http.Request) {
	user := h.getUserFromContext(r)
	userInfo, err := h.UserService.GetUserInfo(user)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}
	err = pkg.WriteJSON(w, http.StatusOK, pkg.Envelope{"user_info": userInfo}, nil)
	if err != nil {
		h.Logger.ErrLog.Print(err)
		w.WriteHeader(500)
	}
}

func (h *Handler) putUserInfo(w http.ResponseWriter, r *http.Request) {
	user := h.getUserFromContext(r)

	var inputUserData struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
	}

	err := json.NewDecoder(r.Body).Decode(&inputUserData)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}

	createdUserInfo := &model.UserInfo{
		Username: inputUserData.Username,
		Phone:    inputUserData.Phone,
		ID:       user.ID,
	}
	newUserInfo, err := h.UserService.ChangeUserInfo(createdUserInfo)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}
	err = pkg.WriteJSON(w, http.StatusOK, pkg.Envelope{"user_info": newUserInfo}, nil)
	if err != nil {
		h.Logger.ErrLog.Print(err)
		w.WriteHeader(500)
	}
}

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := pkg.ReadIDParam(r)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusNotFound, err.Error())
		h.Logger.ErrLog.Print("Not Found")
		return
	}
	userInfo, err := h.UserService.GetUserInfoByID(id)
	if err != nil {
		pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		h.Logger.ErrLog.Print(err.Error())
		return
	}
	err = pkg.WriteJSON(w, http.StatusOK, pkg.Envelope{"user_info": userInfo}, nil)
	if err != nil {
		h.Logger.ErrLog.Print(err)
		w.WriteHeader(500)
	}
}
