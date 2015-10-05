package logging

import (
	l "log"
	"os"

	"github.com/codegangsta/cli"
)

var log *l.Logger
var debug bool = false

func Configure(ctx *cli.Context) {
	debug = ctx.Bool("debug")
	log = l.New(os.Stdout, "", l.Ltime)
}

func doLog(level string, v ...interface{}) {
	log.Print(v...)
}

func doLogf(level string, format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Debug(v ...interface{}) {
	if debug {
		doLog("DEBUG", v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if debug {
		doLogf("DEBUG", format, v...)
	}
}

func Info(v ...interface{}) {
	doLog("INFO", v...)
}

func Infof(format string, v ...interface{}) {
	doLogf("INFO", format, v...)
}

func Error(v ...interface{}) {
	doLog("ERROR", v...)
}

func Errorf(format string, v ...interface{}) {
	doLogf("ERROR", format, v...)
}

func Panic(v ...interface{}) {
	l.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	l.Panicf(format, v...)
}
