package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option struct {
	MaxSizeInMB int // 日志文件大小，单位MB，>=1
	MaxBackups  int // 日志文件最大备份数，>=1
	Compress    bool
	LocalTime   bool
	Formatter   logrus.Formatter
}

var logger *logrus.Logger = New("defalut.log")

func New(filename string, option ...*Option) *logrus.Logger {
	if filename == "" {
		filename = "./unknown.log"
	}
	options := &Option{}
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

	var formatter logrus.Formatter = &defaultFormatter{}
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
		Out:       output,
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	logger.SetReportCaller(true)

	return logger
}

func SetOptions(filename string, option *Option) {
	// lumberjack logger作为logrus的输出
	if option == nil {
		option = &Option{}
	}
	// options
	maxsize := 20
	if option.MaxSizeInMB > 1 {
		maxsize = option.MaxSizeInMB
	}

	maxbackups := 1
	if option.MaxBackups > 1 {
		maxbackups = option.MaxBackups
	}

	var formatter logrus.Formatter = &defaultFormatter{}
	if option.Formatter != nil {
		formatter = option.Formatter
	}
	output := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,    // MB
		MaxBackups: maxbackups, // reserve 1 backup
		Compress:   option.Compress,
		LocalTime:  option.LocalTime,
	}
	logger.SetOutput(output)
	logger.SetFormatter(formatter)
}
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
