package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", nil)

	return router
}
