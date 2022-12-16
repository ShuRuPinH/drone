package utils

import (
	"log"
	"os"
)

func InitLogger(name string) *log.Logger {
	// start logger
	logger := log.Default()
	logger.SetPrefix(name + " ")

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetOutput(file)

	logger.SetFlags(2)

	return logger
}
