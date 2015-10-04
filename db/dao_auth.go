package db

import "time"

type AuthEntry struct {
	Host    string
	Account int
	Realm   string
	Date    *time.Time
}

func (entity AuthEntry) TableName() string {
	return "auth_log"
}

type AuthEntryQuery struct {
	Paging
	Entries  []AuthEntry
	Template AuthEntry
}

func (db *Database) FindAuthEntries(query *AuthEntryQuery) {
	db.engine.
		Model(AuthEntry{}).
		Where(query.Template).
		Count(&query.Total)

	db.engine.
		Model(AuthEntry{}).
		Offset(query.Offset).
		Where(query.Template).
		Limit(query.Limit).
		Order("date DESC").
		Find(&query.Entries)

	query.Count = len(query.Entries)
}
