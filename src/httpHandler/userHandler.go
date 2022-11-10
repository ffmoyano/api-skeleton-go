package httpHandler

import (
	"net/http"
	"notas/src/controller"
	"notas/src/responseWriter"
)

type userHandler struct {
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:
		controller.GetUsers(w, r)
		return
	case r.Method == http.MethodPost:
		controller.InsertUser(w, r)
		return
	default:
		responseWriter.Response(w, http.StatusMethodNotAllowed, response{"Error": "Method " + r.Method + " not allowed."})
		return
	}
}
