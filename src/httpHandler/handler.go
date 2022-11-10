package httpHandler

import (
	"net/http"
	"notas/src/controller"
)

func SetHandlers(mux *http.ServeMux) {
	// TODO - make something to allow same url to be used more than once or declare multiple httpmethods at once

	mux.HandleFunc("/users", Get(controller.GetUsers))
	mux.HandleFunc("/user", Post(controller.InsertUser))
}
