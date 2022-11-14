package reply

import (
	"encoding/json"
	"net/http"
	"notas/src/internal/logger"
)

func Send(writer http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(err.Error())
	}
	writer.WriteHeader(statusCode)
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(response)
	if err != nil {
		logger.Error(err.Error())
	}
}
