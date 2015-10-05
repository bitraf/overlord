package db

import (
	"fmt"

	"github.com/bitraf/overlord/log"
)

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
		log.Debugf(fmt.Sprintf("%s", v[3]), log.Fields{"time": time})
	} else {
		log.Debug(fmt.Sprintf("%s", v[2]))
	}
}

func newLogger() gormLogger {
	return gormLogger{}
}
