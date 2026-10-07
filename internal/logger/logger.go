package logger

import (
	"log"
	"os"
)

var (
	infoLogger = log.New(
		os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime,
	)

	warnLogger = log.New(
		os.Stdout,
		"WARN: ",
		log.Ldate|log.Ltime,
	)

	errorLogger = log.New(
		os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime,
	)
)

func Info(message string) {
	infoLogger.Println(message)
}

func Warn(message string) {
	warnLogger.Println(message)
}

func Error(message string) {
	errorLogger.Println(message)
}