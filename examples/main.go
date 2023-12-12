package main

import (
	"time"

	log "github.com/eric-tech01/simple-log"
	"github.com/sirupsen/logrus"
)

func main() {

	options := &log.Options{}
	logger := log.New("xx_log", options)

	go func() {
		time.Sleep(time.Second * 5)
		logger.Error("set level to error start")
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
