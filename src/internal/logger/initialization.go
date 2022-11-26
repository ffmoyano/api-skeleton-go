package logger

import (
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
	{file: new(os.File), level: 2, log: new(log.Logger), name: "infolog", prefix: "INFO "},
	{file: new(os.File), level: 1, log: new(log.Logger), name: "warnlog", prefix: "WARN "},
	{file: new(os.File), level: 0, log: new(log.Logger), name: "errorlog", prefix: "ERROR "},
	{file: nil, level: 999, log: new(log.Logger), name: "consolelog", prefix: ""},
}
