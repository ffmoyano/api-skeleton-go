package main

import (
	"notas/src/api/server"
	"notas/src/database"
	"notas/src/internal/env"
	"notas/src/internal/logger"
	"os"
)

func init() {
	env.Read(".env")
	logger.OpenLogs()
	database.Open()
}

func main() {

	defer logger.CloseLogs()

	app := server.Get()

	logger.Info("Opening api at port: %s", os.Getenv("port"))
	err := app.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
