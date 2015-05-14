package db

type CheckinQuery struct {
	From    int
	Count   int
	Total   int
	Entries []Checkin
	Desc    bool
}

func (db *Database) FindCheckin(entry *Checkin) bool {
	notFound := db.engine.First(entry).RecordNotFound()
	return !notFound
}

func (db *Database) FindCheckins(query *CheckinQuery) {
	db.engine.Model(Checkin{}).Count(&query.Total)
	db.engine.Offset(query.From).Limit(query.Count).Order("date desc").Find(&query.Entries)
	query.Count = len(query.Entries)
}
