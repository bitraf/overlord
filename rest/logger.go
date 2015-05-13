package rest

import (
	"fmt"
	"time"

	"github.com/bitraf/overlord/log"
	"github.com/labstack/echo"
)

func Logger(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) *echo.HTTPError {
		start := time.Now()
		if he := h(c); he != nil {
			c.Error(he)
		}
		end := time.Now()
		m := c.Request.Method
		p := c.Request.URL.Path
		n := c.Response.Status()

		fields := log.Fields{
			"status": n, "time": end.Sub(start),
		}

		log.Infof(fmt.Sprintf("%s %s", m, p), fields)
		return nil
	}
}
