package route

import (
	"fmt"
	"net/http"

	"notas/src/api/handler"
	"notas/src/internal/service"
)

type UserHandler struct {
}

func (h *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch {
	case request.Method == http.MethodGet:
		getUsers(writer, request)
		return
	case request.Method == http.MethodPost:
		insertUser(writer, request)
		return
	default:
		handler.Response(writer, http.StatusMethodNotAllowed, payload{"Error": "Method " + request.Method + " not allowed."})
		return
	}
}

func getUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := service.GetUsers()
	if err != nil {
		handler.Response(writer, http.StatusInternalServerError,
			payload{"Error": "Ha ocurrido un error al tratar de recuperar los usuarios"})
	} else {
		handler.Response(writer, http.StatusOK, users)
	}
}

func insertUser(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("insert user placeholder")
}
