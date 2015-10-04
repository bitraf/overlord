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

type AccountQuery struct {
	Paging
	Entries  []Account
	Template Account
}

func (db *Database) FindAccount(entry *Account) bool {
	notFound := db.engine.
		Model(Account{}).
		Where(entry).
		First(entry).
		RecordNotFound()

	return !notFound
}

func (db *Database) FindAccounts(query *AccountQuery) {
	db.engine.
		Model(Account{}).
		Where(query.Template).
		Count(&query.Total)

	db.engine.
		Model(Account{}).
		Offset(query.Offset).
		Where(query.Template).
		Limit(query.Limit).
		Order("id ASC").
		Find(&query.Entries)

	query.Count = len(query.Entries)
}
