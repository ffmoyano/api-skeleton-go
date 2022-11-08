package main

import (
	"notas/src/database"
	"notas/src/dotenv"
	"notas/src/logger"
	"notas/src/server"
	"os"
)

func init() {
	dotenv.Read(".env")
	logger.OpenLogs()
	database.Open()
}

func main() {

	defer logger.CloseLogs()

	server := server.Get()

	logger.Info("Opening server at port: %s", os.Getenv("port"))
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Couldn't start server:  %s")
		os.Exit(1)
	}

}
