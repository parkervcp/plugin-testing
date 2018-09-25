package main

import (
	"os"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//Log is a here for the NewLogger
var (
	level   string
	message string

	Log = logrus.New()
)

func setupLogger() {
	if _, err := os.Stat("logs/"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("logs/", 0755)
		} else {
			// other error
		}
	}

	if _, err := os.Stat("./logs/latest.log"); err == nil {
		err := os.Rename("./logs/latest.log", "./logs/"+time.Now().UTC().Format("2006-01-02 15:04")+".log")

		if err != nil {
			Log.Error("failed to move latest logs.", err)
			return
		}
	}

	if _, err := os.Stat("./logs/debug.log"); err == nil {
		err := os.Rename("./logs/debug.log", "./logs/debug-"+time.Now().UTC().Format("2006-01-02 15:04")+".log")

		if err != nil {
			Log.Error("failed to move debug logs.", err)
			return
		}
	}
	Log.Info("Bot logging online")
}

func setLogLevel(level string) {
	if level == "debug" {
		Log.SetLevel(logrus.DebugLevel)
		Log.Debug("log level set to debug")
	} else if level == "info" {
		Log.SetLevel(logrus.InfoLevel)
		Log.Debug("log level set to info")
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "logs/latest.log",
		logrus.ErrorLevel: "logs/latest.log",
		logrus.DebugLevel: "logs/debug.log",
	}
	Log.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

}
