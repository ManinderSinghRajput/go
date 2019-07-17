package mylog


import (
	"bytes"
	"fmt"
	"log"
)


func Info(message string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Info: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Print(&buf)
}

func Error(message string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Error: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Print(&buf)
}

func Debug(message string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Debug: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Print(&buf)
}