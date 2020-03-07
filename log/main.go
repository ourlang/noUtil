package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"time"
)

//建议使用这一种
func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y%m%d%H%M.log",
		//rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	log.SetReportCaller(true) //将函数名和行数放在日志里面
	log.AddHook(lfHook)
}

//切割日志和清理过期日志
func ConfigLocalFilesystemLogger1(filePath string) {
	writer, err := rotatelogs.New(
		filePath+"-%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(filePath),           // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Second*60*3),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatal("Init log failed, err:", err)
	}
	log.SetReportCaller(true) //将函数名和行数放在日志里面
	log.SetOutput(writer)
	log.SetLevel(log.InfoLevel)
}

func main() {
	//ConfigLocalFilesystemLogger1("log")
	ConfigLocalFilesystemLogger("D:/benben", "sentalog", time.Second*60*3, time.Second*60)
	for {
		log.Debug("调试信息")
		log.Info("提示信息")
		log.Warn("警告信息")
		log.Error("错误信息")
		time.Sleep(500 * time.Millisecond)
	}
}
