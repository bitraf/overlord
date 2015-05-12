package log

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func Configure(ctx *cli.Context) {
	formatter := log.TextFormatter{
		TimestampFormat: "15:04:05Z07:00",
		FullTimestamp:   true,
	}
	log.SetFormatter(&formatter)

	level := ctx.String("level")

	if level == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if level == "info" {
		log.SetLevel(log.InfoLevel)
	} else if level == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if level == "error" {
		log.SetLevel(log.ErrorLevel)
	} else if level == "panic" {
		log.SetLevel(log.PanicLevel)
	}
}

type Fields map[string]interface{}

func Debug(message string) {
	log.Debug(message)
}

func Debugf(message string, fields Fields) {
	log.WithFields(log.Fields(fields)).Debug(message)
}

func Info(message string) {
	log.Info(message)
}

func Infof(message string, fields Fields) {
	log.WithFields(log.Fields(fields)).Info(message)
}

func Warn(message string) {
	log.Warn(message)
}

func Warnf(message string, fields Fields) {
	log.WithFields(log.Fields(fields)).Warn(message)
}

func Error(message string) {
	log.Error(message)
}

func Errorf(message string, fields Fields) {
	log.WithFields(log.Fields(fields)).Error(message)
}

func Panic(message string) {
	log.Panic(message)
}

func Panicf(message string, fields Fields) {
	log.WithFields(log.Fields(fields)).Panic(message)
}
