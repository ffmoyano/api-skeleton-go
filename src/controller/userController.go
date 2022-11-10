package controller

import (
	"fmt"
	"net/http"
	"notas/src/responseWriter"
	"notas/src/service"
)

func GetUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := service.GetUsers()
	if err != nil {
		responseWriter.Response(writer, http.StatusInternalServerError, "Ha ocurrido un error al tratar de recuperar los usuarios")
	} else {
		responseWriter.Response(writer, http.StatusOK, users)
	}
}

func InsertUser(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("insert user placeholder")
}
