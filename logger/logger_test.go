package logger

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	myLog := NewLogger("D:/projectLog", "log", time.Second*60*3, time.Second*60, logrus.InfoLevel)
	for {
		myLog.Debug("debug info")
		myLog.Info("prompt message")
		myLog.Warn("warning info")
		myLog.Error("error info")
		time.Sleep(500 * time.Millisecond)
	}
}

func ExampleNewLogger() {
	myLog := NewLogger("./", "logFileName", time.Second*60*3, time.Second*60, logrus.InfoLevel)
	myLog.Error("test error logger write file.")
}
