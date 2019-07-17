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


func Infof(message string, a ...interface{}) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Info: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Printf(buf.String(), a)
}

func Error(message string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Error: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Print(&buf)
}

func Errorf(message string, a ...interface{}) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Error: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Printf(buf.String(), a)
}

func Debug(message string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Debug: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Print(&buf)
}

func Debugf(message string, a ...interface{}) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "Debug: ", log.Lshortfile | log.LstdFlags)
	)

	logger.Print(message)

	fmt.Printf(buf.String(), a)
}