package db

import "github.com/bitraf/overlord/logging"

type logger interface {
	Print(v ...interface{})
}

type gormLogger struct {
}

func (logger gormLogger) Print(v ...interface{}) {
	if len(v) <= 1 {
		return
	}

	level := v[0]

	if level == "sql" {
		time := v[2]
		logging.Debugf("%s (time = %s)", v[3], time)
	} else {
		logging.Debugf("%s", v[2])
	}
}

func newLogger() gormLogger {
	return gormLogger{}
}
