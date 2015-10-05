package rest

import (
	"net/http"
	"strings"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
)

func (server *Server) getUsers(w http.ResponseWriter, r *http.Request) {
	query := db.AccountQuery{}

	query.Offset = parseInt(queryParam(r, "offset"), 0)
	query.Limit = parseInt(queryParam(r, "limit"), 10)
	query.Template = db.Account{}
	query.Template.Type = "user"

	server.db.FindAccounts(&query)

	result := model.ToUsersResult(query)
	serveJson(w, http.StatusOK, result)
}

func (server *Server) getUserById(w http.ResponseWriter, r *http.Request) {
	row := db.Account{}
	row.Type = "user"

	name := pathParam(r, "id")
	if strings.HasPrefix(name, "@") {
		row.Name = name[1:len(name)]
	} else {
		row.Id = parseInt(name, -1)
	}

	has := server.db.FindAccount(&row)
	if !has {
		errorJson(w, http.StatusNotFound, "Could not find user [%s]", name)
		return
	}

	members := db.MemberQuery{}
	members.Limit = 1000
	members.Template = db.Member{}
	members.Template.Account = row.Id

	server.db.FindMembers(&members)

	serveJson(w, http.StatusOK, model.ToUserDetail(row, members.Entries))
}
