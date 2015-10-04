package db

import "time"

type Account struct {
	Id         int `gorm:"primary_key"`
	Name       string
	Type       string
	LastBilled *time.Time `gorm:"column:last_billed"`
}

func (entity Account) TableName() string {
	return "accounts"
}

type Member struct {
	Id           int    `gorm:"primary_key"`
	Name         string `gorm:"column:full_name"`
	Email        string
	Organization string
	Flag         string
	Price        int
	Account      int
	Date         *time.Time
	Recurrence   string
}

func (entity Member) TableName() string {
	return "members"
}

type Paging struct {
	Offset int
	Limit  int
	Count  int
	Total  int
}
