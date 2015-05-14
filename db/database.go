package db

import (
	"fmt"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/log"

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
	url := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", conf.User, conf.Password,
		conf.Addr, conf.Name, conf.Options)

	engine, err := gorm.Open("postgres", url)
	if err != nil {
		log.Panic(err.Error())
	}

	db.engine = &engine
	db.engine.LogMode(conf.ShowSQL)
}
