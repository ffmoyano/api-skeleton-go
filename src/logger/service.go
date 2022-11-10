package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
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

func Error(message string) {
	write(newMessage(message), 0)
}

func Info(message string, args ...any) {
	write(fmt.Sprintf(message, args), 2)
}

func Warn(message string) {
	write(newMessage(message), 1)
}

func write(message string, level int) {

	for _, logger := range loggers {
		if logger.level >= level {
			logger.log.Print(message)
		}
	}
}

func newMessage(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("[%s][%d] : %s", file, line, message)
}
