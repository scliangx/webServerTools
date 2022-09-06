package logger

import (
	"github.com/scliang-strive/webServerTools/config"
	"github.com/scliang-strive/webServerTools/utils"
	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

var LevelMap = map[int]log.Level{}

func init() {
	LevelMap = make(map[int]log.Level)
	LevelMap[0] = log.PanicLevel
	LevelMap[1] = log.ErrorLevel
	LevelMap[2] = log.WarnLevel
	LevelMap[3] = log.InfoLevel
	LevelMap[4] = log.DebugLevel
}

func LogInit() {

	logDir := config.GetConfig().Logger.File
	if !utils.PathExists(logDir) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(err)
		}
	}
	logfile := path.Join(logDir)
	fsWriter, err := rotatelogs.New(
		logfile+"_%Y_%m_%d_%H.log",
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	var multiWriter io.Writer
	stdout := config.GetConfig().Logger.Stdout
	if stdout == "file" {
		multiWriter = io.MultiWriter(fsWriter)
	} else if stdout == "stdout" {
		multiWriter = io.MultiWriter(os.Stdout)
	} else {
		multiWriter = io.MultiWriter(fsWriter, os.Stdout)
	}
	// 显示行号和函数名称
	log.SetReportCaller(true)
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})
	// 日志文件io
	log.SetOutput(multiWriter)
	log.SetLevel(LevelMap[config.GetConfig().Logger.LogLevel])

}
