package httpHandler

import (
	"net/http"
	"notas/src/responseWriter"
)

type response map[string]string

func Get(next http.HandlerFunc) http.HandlerFunc {
	return writeResponse(next, "GET")
}

func Post(next http.HandlerFunc) http.HandlerFunc {
	return writeResponse(next, "POST")
}

func writeResponse(handlerFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == method {
			handlerFunc.ServeHTTP(writer, request)
		} else {
			responseWriter.Response(writer, http.StatusMethodNotAllowed, response{"Error": "Method " + request.Method + " not allowed."})
		}
	}
}
