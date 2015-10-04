package db

import "time"

type Event struct {
	Id      int `gorm:"primary_key"`
	Type    string
	Account int
	Date    *time.Time
}

func (entity Event) TableName() string {
	return "events"
}

type EventQuery struct {
	Paging
	Entries  []Event
	Template Event
}

func (db *Database) FindEvents(query *EventQuery) {
	db.engine.
		Model(Event{}).
		Where(query.Template).
		Count(&query.Total)

	db.engine.
		Model(Event{}).
		Offset(query.Offset).
		Where(query.Template).
		Limit(query.Limit).
		Order("date DESC").
		Find(&query.Entries)

	query.Count = len(query.Entries)
}
