package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

const (
	cookieName = "session"
)

func SetCookie(w http.ResponseWriter) string {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    GetToken(),
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 7),
		MaxAge:   3600,
	}

	http.SetCookie(w, cookie)
	return cookie.Value
}

func GetCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}

func GetToken() string { //TODO: edit error
	token, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return token.String()
}
