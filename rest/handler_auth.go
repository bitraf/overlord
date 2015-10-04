package rest

import (
	"net/http"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func (server *Server) getAuthEntries(c *echo.Context) *echo.HTTPError {
	query := db.AuthEntryQuery{}

	query.Offset = intParam(c, "offset", 0)
	query.Limit = intParam(c, "limit", 10)
	query.Template = db.AuthEntry{}

	if hasParam(c, "byUser") {
		query.Template.Account = intParam(c, "byUser", -1)
	}

	if hasParam(c, "byRealm") {
		query.Template.Realm = stringParam(c, "byRealm", "")
	}

	server.db.FindAuthEntries(&query)

	result := model.ToAuthEntriesResult(query)
	return c.JSON(http.StatusOK, result)
}
