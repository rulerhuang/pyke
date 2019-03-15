package logger

import (
	"io"
	"log"
	"os"
)

var (
	PykeInfo  *log.Logger
	PykeError *log.Logger
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

	PykeInfo = log.New(io.MultiWriter(infoFile, os.Stdout),
		"PykeInfo: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	PykeInfo.Println("PykeInfo logger init")

	PykeError = log.New(io.MultiWriter(errorFile, os.Stdout),
		"PykeError: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	PykeError.Println("PykeError logger init")
}
