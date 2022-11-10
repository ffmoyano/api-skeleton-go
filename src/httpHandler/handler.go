package httpHandler

import (
	"net/http"
)

type response map[string]string

func SetHandlers(mux *http.ServeMux) {
	mux.Handle("/user", &userHandler{})
}
