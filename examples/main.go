package main

import (
	"time"

	logger "github.com/eric-tech01/simple-log"
	"github.com/sirupsen/logrus"
)

func main() {

	logger.SetOptions("new_name.log", &logger.Option{})
	go func() {
		time.Sleep(time.Second * 5)
		logger.Errorf("set level to error start")
		//设置日志级别
		logger.SetLevel(logrus.ErrorLevel)
	}()

	for {

		logger.Debugf("this is debug log")
		logger.Infof("this is info log")
		logger.Warnf("this is warn log")
		logger.Errorf("this is error log")
		time.Sleep(1 * time.Second)
	}
}
