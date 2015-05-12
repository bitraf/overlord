package rest

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func handleStatus(c *echo.Context) *echo.HTTPError {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func StartServer() {
	e := echo.New()
	e.Use(mw.Logger)

	e.Get("/", handleStatus)

	e.Run(":1323")
}
