package rest

import (
	"fmt"
	"net/http"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func (server *Server) getMembers(c *echo.Context) *echo.HTTPError {
	query := db.MemberQuery{}

	query.Offset = intParam(c, "offset", 0)
	query.Limit = intParam(c, "limit", 10)
	query.Template = db.Member{}

	if hasParam(c, "byUser") {
		query.Template.Account = intParam(c, "byUser", -1)
	}

	server.db.FindMembers(&query)

	result := model.ToMembersResult(query)
	return c.JSON(http.StatusOK, result)
}

func (server *Server) getMemberById(c *echo.Context) *echo.HTTPError {
	row := db.Member{}
	row.Id = intParam(c, "id", -1)

	has := server.db.FindMember(&row)
	if !has {
		return errorJson(c, http.StatusNotFound,
			fmt.Sprintf("Could not find member [%v]", row.Id))
	}

	return c.JSON(
		http.StatusOK,
		model.ToMemberDetail(row),
	)
}
