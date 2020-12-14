package logger

import (
	"log"
)

//func Log(level, message string, args ...interface{}) {
//	params := append([]interface{}{level + ": ", message + " "}, args...)
//
//	log.Print(params...)
//}

func Info(message string, args ...interface{}) {
	Log("info", message, args...)
}

func Error(message string, args ...interface{}) {
	Log("error", message, args...)
}

func Log(level, message string, args ...interface{}) {
	params := append([]interface{}{level + ": ", message + " "}, sanitize(args...)...)

	log.Print(params...)
}
