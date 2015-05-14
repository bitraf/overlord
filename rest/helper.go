package rest

import (
	"strconv"

	"github.com/labstack/echo"
)

func stringParam(c *echo.Context, name string, def string) string {
	str := c.Param(name)

	if len(str) == 0 {
		str = c.Request.URL.Query().Get(name)
	}

	if len(str) == 0 {
		return def
	}

	return str
}

func intParam(c *echo.Context, name string, def int) int {
	str := stringParam(c, name, "")

	if len(str) == 0 {
		return def
	}

	val, err := strconv.Atoi(str)

	if err != nil {
		return def
	}

	return val
}
