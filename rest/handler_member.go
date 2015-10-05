package rest

import (
	"net/http"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
)

func (server *Server) getMembers(w http.ResponseWriter, r *http.Request) {
	query := db.MemberQuery{}

	query.Offset = parseInt(queryParam(r, "offset"), 0)
	query.Limit = parseInt(queryParam(r, "limit"), 10)
	query.Template = db.Member{}

	byUser := queryParam(r, "byUser")
	if len(byUser) > 0 {
		query.Template.Account = parseInt(byUser, -1)
	}

	server.db.FindMembers(&query)

	result := model.ToMembersResult(query)
	serveJson(w, http.StatusOK, result)
}

func (server *Server) getMemberById(w http.ResponseWriter, r *http.Request) {
	row := db.Member{}
	row.Id = parseInt(pathParam(r, "id"), -1)

	has := server.db.FindMember(&row)
	if !has {
		errorJson(w, http.StatusNotFound, "Could not find member [%v]", row.Id)
		return
	}

	serveJson(w, http.StatusOK, model.ToMemberDetail(row))
}
