package analytics

import (
	"log"
	"os"
	"path"
)

const (
	ANALYTICS_DIR string = ".freighter"
	LOG_FILE      string = "logs.txt"
)

type Analytics struct {
	LogFile *os.File
}

func NewAnalytics() *Analytics {
	if _, err := os.Stat(ANALYTICS_DIR); os.IsNotExist(err) {
		err := os.MkdirAll(ANALYTICS_DIR, 0755)
		if err != nil {
			log.Fatalf("Could not create directory: %s", ANALYTICS_DIR)
		}
	}

	logFile, err := os.OpenFile(path.Join(ANALYTICS_DIR, LOG_FILE), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Could not open log-file: %s", LOG_FILE)
	}

	return &Analytics{
		LogFile: logFile,
	}
}
