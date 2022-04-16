package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type MyLog struct {
	Logger *logrus.Logger
}

func (l *MyLog) Init(logFolder string) {
	var log = logrus.New()
	logFile := logFolder + "/" + string(time.Now().Format("20060102150405")) + ".log"
	var file, err = os.Create(logFile)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
		os.Exit(-1)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	log.SetFormatter(&logrus.JSONFormatter{})

	l.Logger = log
}

func (l *MyLog) LogInfo(f string, what ...interface{}) {
	l.Logger.Infof(f, what)
}

func (l *MyLog) LogWarn(f string, what ...interface{}) {
	l.Logger.Warnf(f, what)
}

func (l *MyLog) LogFatal(f string, what ...interface{}) {
	l.Logger.Errorf(f, what)
	os.Exit(-1)
}
