package rest

import (
	"net/http"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func (server *Server) getEvents(c *echo.Context) *echo.HTTPError {
	query := db.EventQuery{}

	query.Offset = intParam(c, "offset", 0)
	query.Limit = intParam(c, "limit", 10)
	query.Template = db.Event{}

	if hasParam(c, "byUser") {
		query.Template.Account = intParam(c, "byUser", -1)
	}

	if hasParam(c, "byType") {
		query.Template.Type = stringParam(c, "byType", "")
	}

	server.db.FindEvents(&query)

	result := model.ToEventsResult(query)
	return c.JSON(http.StatusOK, result)
}
