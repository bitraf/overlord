package model

import (
	"time"

	"github.com/bitraf/overlord/db"
)

type AuthEntry struct {
	Host  string     `json:"host"`
	Date  *time.Time `json:"date"`
	Realm string     `json:"realm"`
	User  int        `json:"user"`
}

func ToAuthEntry(row *db.AuthEntry) AuthEntry {
	return AuthEntry{
		Host:  row.Host,
		Date:  row.Date,
		Realm: row.Realm,
		User:  row.Account,
	}
}

func ToAuthEntries(rows []db.AuthEntry) []AuthEntry {
	result := make([]AuthEntry, len(rows))

	for index := range rows {
		result[index] = ToAuthEntry(&rows[index])
	}

	return result
}

func ToAuthEntriesResult(query db.AuthEntryQuery) PagedResult {
	return PagedResult{
		Offset: query.Offset,
		Limit:  query.Limit,
		Count:  query.Count,
		Total:  query.Total,
		Result: ToAuthEntries(query.Entries),
	}
}
