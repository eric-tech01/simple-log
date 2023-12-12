package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options struct {
	MaxSizeInMB int // 日志文件大小，单位MB，>=1
	MaxBackups  int // 日志文件最大备份数，>=1
	Formatter   logrus.Formatter
}

func New(filename string, option ...*Options) *logrus.Logger {
	if filename == "" {
		filename = "./unknown.log"
	}
	options := &Options{}
	if len(option) > 0 {
		options = option[0]
	}

	// options
	maxsize := 20
	if options.MaxSizeInMB > 1 {
		maxsize = options.MaxSizeInMB
	}

	maxbackups := 1
	if options.MaxBackups > 1 {
		maxbackups = options.MaxBackups
	}

	var formatter logrus.Formatter = &myFormatter{}
	if options.Formatter != nil {
		formatter = options.Formatter
	}

	// lumberjack logger作为logrus的输出
	output := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,    // MB
		MaxBackups: maxbackups, // reserve 1 backup
		Compress:   true,
		LocalTime:  true,
	}

	logger := &logrus.Logger{
		Out: output,
		// Formatter: &logrus.TextFormatter{},
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	logger.SetReportCaller(true)

	return logger
}
