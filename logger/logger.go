package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

const DefaultInfoLogPath = "./info_logs.log"
const DefaultErrorLogPath = "./error_logs.log"

func init() {
	infoFile, err := os.OpenFile(DefaultInfoLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open info log:", err)
	}

	errorFile, err := os.OpenFile(DefaultErrorLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open error log:", err)
	}

	Info = log.New(io.MultiWriter(infoFile, os.Stdout),
		"Info: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Info.Println("Info logger init")

	Error = log.New(io.MultiWriter(errorFile, os.Stdout),
		"Error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error.Println("Error logger init")
}
