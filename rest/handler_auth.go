package rest

import (
	"net/http"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
)

func (server *Server) getAuthEntries(w http.ResponseWriter, r *http.Request) {
	query := db.AuthEntryQuery{}

	query.Offset = parseInt(queryParam(r, "offset"), 0)
	query.Limit = parseInt(queryParam(r, "limit"), 10)
	query.Template = db.AuthEntry{}

	byUser := queryParam(r, "byUser")
	if len(byUser) > 0 {
		query.Template.Account = parseInt(byUser, -1)
	}

	byRealm := queryParam(r, "byRealm")
	if len(byRealm) > 0 {
		query.Template.Realm = byRealm
	}

	server.db.FindAuthEntries(&query)

	result := model.ToAuthEntriesResult(query)
	serveJson(w, http.StatusOK, result)
}
