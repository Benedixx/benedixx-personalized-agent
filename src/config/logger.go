package config

import (
	"os"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&formatter.Formatter{
		NoColors:    false,
		HideKeys:    true,
		FieldsOrder: []string{"time", "level", "msg"},
	})
	Log.Info("Logger initialized with nested formatter")
	Log.WithField("component", "api").Info("API request received")
}

func Info(msg string, fields ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"fields": fields,
	}).Info(msg)
}

func Warn(msg string, fields ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"fields": fields,
	}).Warn(msg)
}

func Error(msg string, fields ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"fields": fields,
	}).Error(msg)
}

func Critical(msg string, fields ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"fields": fields,
	}).Fatal(msg)
}
