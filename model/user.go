package model

import (
	"time"

	"github.com/bitraf/overlord/db"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	*UserDetail
}

type UserDetail struct {
	LastBilled *time.Time `json:"lastBilled"`
	Members    []int      `json:"members"`
}

func ToUser(row db.Account) User {
	return User{
		Id:   row.Id,
		Name: row.Name,
	}
}

func ToUserDetail(row db.Account, members []db.Member) User {
	result := ToUser(row)
	details := UserDetail{}
	result.UserDetail = &details

	details.LastBilled = row.LastBilled
	details.Members = make([]int, len(members))

	for index := range members {
		details.Members[index] = members[index].Id
	}

	return result
}

func ToUsers(rows []db.Account) []User {
	result := make([]User, len(rows))

	for index := range rows {
		result[index] = ToUser(rows[index])
	}

	return result
}

func ToUsersResult(query db.AccountQuery) PagedResult {
	return PagedResult{
		Offset: query.Offset,
		Limit:  query.Limit,
		Count:  query.Count,
		Total:  query.Total,
		Result: ToUsers(query.Entries),
	}
}
