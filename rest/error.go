package rest

import (
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func (server *Server) error(c *echo.Context, e model.Error) *echo.HTTPError {
	return c.JSON(e.Code, e)
}
