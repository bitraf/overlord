package db

import (
	"github.com/bitraf/overlord/log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Database struct {
	Val string
}

func New() Database {
	return Database{}
}

func (db *Database) Open() {
	handle, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	if err != nil {
		log.Panic(err.Error())
	}

	handle.DB()
}

func (db *Database) FindMembers() {
	println(db.Val)
}
