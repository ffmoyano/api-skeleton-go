package logger

import (
	"fmt"
	"log"
	"os"
)

var err error

func OpenLogs() {
	if _, err = os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, logger := range loggers {
		logger = logger.Open()
		loggers[i] = logger
	}
}

func CloseLogs() {
	for _, logger := range loggers {
		logger.Close()
	}
}

func Error(message string, args ...any) {
	write(fmt.Sprintf(message, args), 0)
}

func Info(message string, args ...any) {
	write(fmt.Sprintf(message, args), 2)
}

func Warn(message string, args ...any) {
	write(fmt.Sprintf(message, args), 1)
}

func write(message string, level int) {

	for _, logger := range loggers {
		if logger.level >= level {
			logger.log.Print(message)
		}
	}
}
