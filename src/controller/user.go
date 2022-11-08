package controller

import (
	"fmt"
	"net/http"
	"notas/src/logger"
	"notas/src/response"
	"notas/src/service"
)

func GetUsers(writer http.ResponseWriter, _ *http.Request) {
	users, err := service.GetUsers()
	if err != nil {
		logger.Error("Error parsing to response the users retrieved from database: %s", err.Error())
		response.Response(writer, http.StatusBadRequest, fmt.Sprintf("Error parsing to response the users from database: %s", err))
	} else {
		response.Response(writer, http.StatusOK, users)
	}
}
