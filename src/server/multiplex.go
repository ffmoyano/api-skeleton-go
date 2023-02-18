package server

import (
	"net/http"
	"rest-skeleton-go/src/handler"
)

func SetHandlers(mux *http.ServeMux) {
	mux.Handle("/user", &handler.UserHandler{})
}
