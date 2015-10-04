package rest

import (
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func errorJson(c *echo.Context, code int, message string) *echo.HTTPError {
	return c.JSON(code, model.Error{Code: code, Message: message})
}
