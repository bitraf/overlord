package rest

import "github.com/labstack/echo"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (server *Server) error(c *echo.Context, code int, message string) *echo.HTTPError {
	result := Error{}
	result.Code = code
	result.Message = message
	return c.JSON(code, result)
}
