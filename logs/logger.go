package logs

import (
	"os"

	"github.com/labstack/gommon/log"
)

func NewLogger(writeToFile bool) *log.Logger {
	logger := log.New("note")
	if writeToFile {
		logFile, err := os.OpenFile("app.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		logger.SetOutput(logFile)
	}
	logger.SetLevel(log.INFO)
	logger.SetHeader("${time_rfc3339} ${level} ${short_time}:${line} ${message}")

	logger.Info("Started")
	return logger
}
