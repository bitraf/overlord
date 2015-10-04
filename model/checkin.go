package model

import (
	"time"

	"github.com/bitraf/overlord/db"
)

type Checkin struct {
	Id   int        `json:"id"`
	Date *time.Time `json:"date"`
	Type string     `json:"type"`
	User int        `json:"user"`
}

func ToCheckin(row *db.Checkin) Checkin {
	return Checkin{
		Id:   row.Id,
		Date: row.Date,
		Type: row.Type,
		User: row.Account,
	}
}

func ToCheckins(rows []db.Checkin) []Checkin {
	result := make([]Checkin, len(rows))

	for index := range rows {
		result[index] = ToCheckin(&rows[index])
	}

	return result
}

func ToCheckinsResult(query db.CheckinQuery) PagedResult {
	return PagedResult{
		Offset: query.Offset,
		Limit:  query.Limit,
		Count:  query.Count,
		Total:  query.Total,
		Result: ToCheckins(query.Entries),
	}
}
