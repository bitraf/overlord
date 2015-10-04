package model

import (
	"time"

	"github.com/bitraf/overlord/db"
)

type Event struct {
	Id   int        `json:"id"`
	Date *time.Time `json:"date"`
	Type string     `json:"type"`
	User int        `json:"user"`
}

func ToEvent(row *db.Event) Event {
	return Event{
		Id:   row.Id,
		Date: row.Date,
		Type: row.Type,
		User: row.Account,
	}
}

func ToEvents(rows []db.Event) []Event {
	result := make([]Event, len(rows))

	for index := range rows {
		result[index] = ToEvent(&rows[index])
	}

	return result
}

func ToEventsResult(query db.EventQuery) PagedResult {
	return PagedResult{
		Offset: query.Offset,
		Limit:  query.Limit,
		Count:  query.Count,
		Total:  query.Total,
		Result: ToEvents(query.Entries),
	}
}
