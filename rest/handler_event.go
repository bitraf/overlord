package rest

import (
	"net/http"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
)

func (server *Server) getEvents(w http.ResponseWriter, r *http.Request) {
	query := db.EventQuery{}

	query.Offset = parseInt(queryParam(r, "offset"), 0)
	query.Limit = parseInt(queryParam(r, "limit"), 10)
	query.Template = db.Event{}

	byUser := queryParam(r, "byUser")
	if len(byUser) > 0 {
		query.Template.Account = parseInt(byUser, -1)
	}

	byType := queryParam(r, "byType")
	if len(byType) > 0 {
		query.Template.Type = byType
	}

	server.db.FindEvents(&query)

	result := model.ToEventsResult(query)
	serveJson(w, http.StatusOK, result)
}
