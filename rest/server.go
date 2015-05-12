package rest

import (
	"net/http"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/log"
	"github.com/labstack/echo"
)

func handleStatus(c *echo.Context) *echo.HTTPError {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func StartServer() {
	e := echo.New()
	e.Use(Logger)

	e.Get("/", handleStatus)

	addr := config.C.Server.Addr
	log.Infof("Running server.", log.Fields{"addr": addr})
	e.Run(addr)
}
