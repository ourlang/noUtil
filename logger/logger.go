// Copyright 2020 The ourlang Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// package logger is log processing framework tool
// It can be called directly or set through the configuration file
package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

const (
	DefaultLogPath     = "./"
	DefaultLogFileName = "go"
)

// Creates a new logger.
// `logPath` is logger file path ,default is the root path of the current project.
// `logFileName` is logger file name
// `maxAge` maximum file save time.
// `rotationTime` log cut interval
// `logLevel` logger level
// usage:
//   log:=NewLogger("D:/projectLog", "log", time.Second*60*3, time.Second*60, logrus.InfoLevel)
//   log.Info("https://github.com/ourlang/noUtil")
//   It's recommended to make this a global instance called `log`.
func NewLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration, logLevel logrus.Level) *logrus.Logger {
	myLog := logrus.New()
	if len(logPath) == 0 {
		logPath = DefaultLogPath
	}
	if len(logFileName) == 0 {
		logFileName = DefaultLogFileName
	}
	if maxAge <= 0 {
		maxAge = 24 * time.Hour
	}
	if rotationTime <= 0 {
		maxAge = 30 * time.Minute
	}
	if logLevel < 0 || logLevel > 6 {
		logLevel = logrus.ErrorLevel
	}
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		myLog.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	//set different output purposes for different levels
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	myLog.SetLevel(logLevel)
	//put the function name and the number of lines in the log
	myLog.SetReportCaller(true)
	myLog.AddHook(lfHook)
	return myLog
}
