package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var timeFormat = time.Now().Format("02-01-2006")

type logger struct {
	file   *os.File
	level  int
	log    *log.Logger
	name   string
	prefix string
}

// when we call a log of level x all logs with level >= x are also called
var loggers = []logger{
	logger{file: new(os.File), level: 2, log: new(log.Logger), name: "infolog", prefix: "INFO "},
	logger{file: new(os.File), level: 1, log: new(log.Logger), name: "warnlog", prefix: "WARN "},
	logger{file: new(os.File), level: 0, log: new(log.Logger), name: "errorlog", prefix: "ERROR "},
	logger{file: nil, level: 999, log: new(log.Logger), name: "consolelog", prefix: ""},
}

func (logger logger) Close() {
	err = logger.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (logger logger) Open() logger {
	if logger.file != nil {

		logger.file, err = os.OpenFile(fmt.Sprintf("logs/%s_%s.log", logger.name, timeFormat),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		logger.log = log.New(logger.file, logger.prefix, log.Ldate|log.Ltime|log.Lshortfile)

	} else { // console log has no file associated
		logger.log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return logger
}
