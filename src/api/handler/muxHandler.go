package handler

import (
	"net/http"
)

func SetHandlers(mux *http.ServeMux) {
	mux.Handle("/user", &UserHandler{})
}
