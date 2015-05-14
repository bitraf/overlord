package db

import "time"

type Account struct {
	Id         int `gorm:"primary_key"`
	Name       string
	Type       string    // product, user, p2k12
	LastBilled time.Time // only 1 row. probably not used.
}

func (a *Account) TableName() string {
	return "accounts"
}

type Checkin struct {
	Id      int `gorm:"primary_key"`
	Account Account
	Date    time.Time
	Type    string // checkin, checkout
}

func (c *Checkin) TableName() string {
	return "checkins"
}

type AuthLog struct {
	Host      string
	Account   int
	Timestamp time.Time
	Realm     string
}

type AccountAlias struct {
	Account int
	Alias   string
}

type Member struct {
	Id           int
	Timestamp    time.Time
	FullName     string
	Email        string
	Account      int
	Organization string
	Price        float32
	Recurrence   string // 1 mon
	Flag         string
}

type Event struct {
	Id        int
	Timestamp time.Time
	Type      string // inform-door-uri
	Account   int
	Amount    float32 // no rows. probably not used.
}
