package db

import "time"

/*
func (db *Database) FindCheckin(entry *Checkin) bool {
	notFound := db.engine.First(entry).RecordNotFound()
	return !notFound
}
*/

type Checkin struct {
	Id      int `gorm:"primary_key"`
	Type    string
	Account int
	Date    *time.Time
}

func (entity Checkin) TableName() string {
	return "checkins"
}

type CheckinQuery struct {
	Paging
	Entries  []Checkin
	Template Checkin
}

func (db *Database) FindCheckins(query *CheckinQuery) {
	db.engine.
		Model(Checkin{}).
		Where(query.Template).
		Count(&query.Total)

	db.engine.
		Model(Checkin{}).
		Offset(query.Offset).
		Where(query.Template).
		Limit(query.Limit).
		Order("date DESC").
		Find(&query.Entries)

	query.Count = len(query.Entries)
}
