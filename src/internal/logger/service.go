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
		logger = logger.open()
		loggers[i] = logger
	}
}

func CloseLogs() {
	for _, logger := range loggers {
		logger.close()
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

func (logger logger) close() {
	err = logger.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (logger logger) open() logger {
	if logger.file != nil {

		logger.file, err = os.OpenFile(fmt.Sprintf("logs/%s_%s.log", logger.name, timeFormat),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		logger.log = log.New(logger.file, logger.prefix, log.Ldate|log.Ltime)

	} else { // console log has no file associated
		logger.log = log.New(os.Stdout, "", log.LstdFlags)
	}
	return logger
}

func write(message string, level int) {
	var prefix string
	for _, logger := range loggers {
		if logger.level >= level {
			switch level {
			case 0:
				prefix = "[ERROR]"
			case 1:
				prefix = "[WARN]"
			case 2:
				prefix = "[INFO]"
			}
			logger.log.Print(prefix + " " + message)
		}
	}
}

func newMessage(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("[%s][%d] : %s", file, line, message)
}
