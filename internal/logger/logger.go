package logger

import (
	"log"
	"os"
)

type Logger struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func NewLogger() Logger {
	newLogger := Logger{
		InfoLog: log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile),
		ErrLog: log.New(os.Stdout, "ERROR: ", log.Ldate | log.Ltime | log.Lshortfile),
	}
	return newLogger
}
