package handler

import (
	"net/http"
	"notas/src/api/route"
)

func SetHandlers(mux *http.ServeMux) {
	mux.Handle("/user", &route.UserHandler{})
}
