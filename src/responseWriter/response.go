package responseWriter

import (
	"encoding/json"
	"net/http"
	"notas/src/logger"
)

// Response sets headers and sends the responseWriter with the result of the dbPool query as a json
func Response(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(err.Error())
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		logger.Error(err.Error())
	}
}
