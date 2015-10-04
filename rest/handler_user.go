package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func (server *Server) getUsers(c *echo.Context) *echo.HTTPError {
	query := db.AccountQuery{}

	query.Offset = intParam(c, "offset", 0)
	query.Limit = intParam(c, "limit", 10)
	query.Template = db.Account{}
	query.Template.Type = "user"

	server.db.FindAccounts(&query)

	result := model.ToUsersResult(query)
	return c.JSON(http.StatusOK, result)
}

func (server *Server) getUserById(c *echo.Context) *echo.HTTPError {
	row := db.Account{}
	row.Type = "user"

	name := stringParam(c, "id", "")
	if strings.HasPrefix(name, "@") {
		row.Name = name[1:len(name)]
	} else {
		row.Id = intParam(c, "id", -1)
	}

	has := server.db.FindAccount(&row)
	if !has {
		return errorJson(c, http.StatusNotFound,
			fmt.Sprintf("Could not find user [%s]", name))
	}

	members := db.MemberQuery{}
	members.Limit = 1000
	members.Template = db.Member{}
	members.Template.Account = row.Id

	server.db.FindMembers(&members)

	return c.JSON(
		http.StatusOK,
		model.ToUserDetail(row, members.Entries),
	)
}
