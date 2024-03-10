package handler

import (
	"context"
	"net/http"

	"github.com/mystpen/test-task-Mobydev/internal/model"
	"github.com/mystpen/test-task-Mobydev/pkg"
)

type contextKey string

const ctxKey contextKey = "user"

func (h *Handler) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := pkg.GetCookie(r)
		user := &model.User{}
		switch err {
		case http.ErrNoCookie:
		case nil:
			user, err = h.UserService.GetUserByToken(cookie.Value)
			if err != nil {
				pkg.DeleteCookie(w)
				h.Logger.ErrLog.Print(err.Error())
				pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
				return
			}
		default:
			pkg.ErrorResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), ctxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := h.getUserFromContext(r)
		if len(user.Username) == 0 {
			message := "Unauthorized"
			h.Logger.ErrLog.Print(message)
			pkg.ErrorResponse(w, r, http.StatusUnauthorized, message)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) getUserFromContext(r *http.Request) *model.User {
	user, ok := r.Context().Value(ctxKey).(*model.User)

	if !ok {
		message := "No user data"
		h.Logger.ErrLog.Print(message)
		return nil
	}
	return user
}
