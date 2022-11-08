package handler

import (
	"net/http"
	"notas/src/controller"
)

func SetHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/users", Get(controller.GetUsers))
}
