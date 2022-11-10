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

	app := server.Get()

	logger.Info("Opening server at port: %s", os.Getenv("port"))
	err := app.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
