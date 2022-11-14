package server

import (
	"net/http"
	"notas/src/api/handler"
)

func SetHandlers(mux *http.ServeMux) {
	mux.Handle("/user", &handler.UserHandler{})
}
