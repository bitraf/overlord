package rest

import (
	"net/http"
	"time"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/model"
	"github.com/labstack/echo"
)

func (server *Server) getStatus(c *echo.Context) *echo.HTTPError {
	now := time.Now()
	duration := now.Sub(server.startTime).String()

	res := model.Status{
		Version: config.C.Version,
		Uptime:  duration,
	}

	return c.JSON(http.StatusOK, res)
}
