package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/INebotov/JustNets/backend/config"
	"github.com/sirupsen/logrus"
)

type MyLog struct {
	Logger *logrus.Logger
	Prefix string
}

func (l *MyLog) Init(logFolder string, Prefix string) {
	l.Prefix = Prefix
	var log = logrus.New()
	logFile := fmt.Sprintf(config.GetLogPath(), logFolder, l.Prefix, string(time.Now().Format("20060102150405")))
	var file, err = os.Create(logFile)
	if err != nil {
		fmt.Println("Prefix - Could Not Open Log File : " + err.Error())
		os.Exit(-1)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	log.SetFormatter(&logrus.JSONFormatter{})

	l.Logger = log
}

func (l *MyLog) LogInfo(f string, what ...interface{}) {
	l.Logger.WithField("Prefix", l.Prefix).Infof(f, what)
}

func (l *MyLog) LogWarn(f string, what ...interface{}) {
	l.Logger.WithField("Prefix", l.Prefix).Warnf(f, what)
}

func (l *MyLog) LogFatal(f string, what ...interface{}) {
	l.Logger.WithField("Prefix", l.Prefix).Errorf(f, what)
	os.Exit(-1)
}
