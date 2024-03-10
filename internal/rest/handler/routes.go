package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/signup", h.signup)
	router.HandlerFunc(http.MethodPost, "/signin", h.signin)
	router.HandlerFunc(http.MethodGet, "/user/info",  h.requireAuth(h.getUserInfo))
	router.HandlerFunc(http.MethodPut, "/user/info", h.requireAuth(h.putUserInfo))
	router.HandlerFunc(http.MethodGet, "/users/:id", h.getUserByID)
	router.HandlerFunc(http.MethodGet, "/videos/:id", h.getVideoInfo) 
	router.HandlerFunc(http.MethodPut, "/videos/update/:id", h.requireAdmin(h.putVideoInfo)) //only admin

	return h.middleware(router)
}
