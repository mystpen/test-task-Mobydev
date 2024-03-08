package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/signup", h.signup)
	router.HandlerFunc(http.MethodPost, "/signin", h.signin)
	router.HandlerFunc(http.MethodGet, "/user/info", nil)
	router.HandlerFunc(http.MethodPut, "/user/info", nil)
	router.HandlerFunc(http.MethodGet, "/users/:id", nil)
	// router.HandlerFunc(http.MethodPost, "/video/create", nil)
	router.HandlerFunc(http.MethodPut, "/videos/update", nil) //only admin

	return router
}
