package logger

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewFilesystemLogger(t *testing.T) {
	ExampleNewFilesystemLogger()
}

func ExampleNewFilesystemLogger() {
	//ConfigLocalFilesystemLogger1("log")
	NewFilesystemLogger("D:/projectLog", "log", time.Second*60*3, time.Second*60, log.InfoLevel)
	for {
		log.Debug("debug info")
		log.Info("prompt message")
		log.Warn("warning info")
		log.Error("error info")
		time.Sleep(500 * time.Millisecond)
	}
}

func TestFilesystemLogger(t *testing.T) {
	ExampleFilesystemLogger()
}
func ExampleFilesystemLogger() {
	FilesystemLogger("log", time.Second*60*3, time.Second*60, log.InfoLevel)
	for {
		log.Debug("debug info")
		log.Info("prompt message")
		log.Warn("warning info")
		log.Error("error info")
		time.Sleep(500 * time.Millisecond)
	}
}
