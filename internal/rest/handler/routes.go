package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/signup", h.signup)
	router.HandlerFunc(http.MethodPost, "/signin", h.signin)
	router.HandlerFunc(http.MethodGet, "/user/info", h.getUserInfo)
	router.HandlerFunc(http.MethodPut, "/user/info", h.putUserInfo)
	router.HandlerFunc(http.MethodGet, "/users/:id", h.getUserByID)
	// router.HandlerFunc(http.MethodPost, "/video/create", nil)
	router.HandlerFunc(http.MethodPut, "/videos/update", nil) //only admin
	router.HandlerFunc(http.MethodPost, "/videos/create", nil) //only admin

	return router
}
