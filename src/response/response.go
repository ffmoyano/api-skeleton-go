package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notas/src/logger"
)

// Response sets headers and sends the response with the result of the dbPool query as a json
func Response(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		logger.Error(fmt.Sprintf("Couldn't write response:  %s", err))
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
	if err != nil {
		logger.Error(fmt.Sprintf("Couldn't write response:  %s", err))
	}
}
