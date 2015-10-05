package db

import (
	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/logging"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Database struct {
	engine *gorm.DB
}

func New() Database {
	return Database{}
}

func (db *Database) Open() {
	conf := config.C.Database
	engine, err := gorm.Open(conf.Driver, conf.Url)
	if err != nil {
		logging.Panic(err.Error())
	}

	db.engine = &engine
	db.engine.LogMode(conf.ShowSQL)
	db.engine.SetLogger(newLogger())
}
