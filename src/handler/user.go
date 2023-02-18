package handler

import (
	"encoding/json"
	"github.com/ffmoyano/gofer/logger"
	"net/http"
	"rest-skeleton-go/src/database"
	"rest-skeleton-go/src/entity"
	"rest-skeleton-go/src/reply"
)

type UserHandler struct {
}

func (handler *UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch {
	case request.Method == http.MethodGet:
		handler.get(writer, request)
		return
	case request.Method == http.MethodPost:
		handler.insert(writer, request)
		return
	default:
		reply.Send(writer, http.StatusMethodNotAllowed, map[string]string{"Error": "Method " + request.Method + " not allowed."})
		return
	}
}

func (handler *UserHandler) get(writer http.ResponseWriter, _ *http.Request) {
	users, err := database.GetUsers()
	if err != nil {
		reply.Send(writer, http.StatusInternalServerError,
			map[string]string{"Error": "Ha ocurrido un error al tratar de recuperar los usuarios"})
	} else {
		reply.Send(writer, http.StatusOK, users)
	}
}

func (handler *UserHandler) insert(writer http.ResponseWriter, request *http.Request) {

	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		logger.Error(err.Error())
		reply.Send(writer, http.StatusInternalServerError,
			map[string]string{"Error": "Ha ocurrido un error al tratar de insertar el usuario"})
	}

}
