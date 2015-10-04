package model

import (
	"time"

	"github.com/bitraf/overlord/db"
)

type Member struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	User int    `json:"user"`
	*MemberDetail
}

type MemberDetail struct {
	Email        string     `json:"email"`
	Organization string     `json:"organization"`
	MemberSince  *time.Time `json:"memberSince"`
	Price        int        `json:"price"`
	Flag         string     `json:"flag"`
	Recurrence   string     `json:"recurrence"`
}

func ToMember(row db.Member) Member {
	return Member{
		Id:   row.Id,
		Name: row.Name,
		User: row.Account,
	}
}

func ToMembers(rows []db.Member) []Member {
	result := make([]Member, len(rows))

	for index := range rows {
		result[index] = ToMember(rows[index])
	}

	return result
}

func ToMemberDetail(row db.Member) Member {
	result := ToMember(row)

	details := MemberDetail{
		Price:        row.Price,
		Organization: row.Organization,
		Email:        row.Email,
		MemberSince:  row.Date,
		Flag:         row.Flag,
		Recurrence:   row.Recurrence,
	}

	result.MemberDetail = &details
	return result
}

func ToMembersResult(query db.MemberQuery) PagedResult {
	return PagedResult{
		Offset: query.Offset,
		Limit:  query.Limit,
		Count:  query.Count,
		Total:  query.Total,
		Result: ToMembers(query.Entries),
	}
}
