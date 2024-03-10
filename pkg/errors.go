package pkg

import (
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := Envelope{"error": message}

	err := WriteJSON(w, status, env, nil)
	if err != nil{
		log.Print(err) //TODO: logs
		w.WriteHeader(500)
	}
}
