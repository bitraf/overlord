package db

import "time"

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

type MemberQuery struct {
	Paging
	Entries  []Member
	Template Member
}

func (db *Database) FindMember(entry *Member) bool {
	notFound := db.engine.
		Model(Member{}).
		Where(entry).
		First(entry).
		RecordNotFound()

	return !notFound
}

func (db *Database) FindMembers(query *MemberQuery) {
	db.engine.
		Model(Member{}).
		Where(query.Template).
		Count(&query.Total)

	db.engine.
		Model(Member{}).
		Offset(query.Offset).
		Where(query.Template).
		Limit(query.Limit).
		Order("id ASC").
		Find(&query.Entries)

	query.Count = len(query.Entries)
}
