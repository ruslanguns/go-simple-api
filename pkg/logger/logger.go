package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func New(w io.Writer) *Logger {
	if w == nil {
		w = os.Stdout
	}
	return &Logger{
		infoLog:  log.New(w, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(w, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLog.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLog.Printf(format, v...)
}
