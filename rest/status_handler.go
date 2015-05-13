package rest

import (
	"net/http"
	"time"

	"github.com/bitraf/overlord/config"
	"github.com/labstack/echo"
)

type Status struct {
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

func (server *Server) handleStatus(c *echo.Context) *echo.HTTPError {
	now := time.Now()
	duration := now.Sub(server.startTime).String()

	res := Status{
		Version: config.C.Version,
		Uptime:  duration,
	}

	return c.JSON(http.StatusOK, res)
}
