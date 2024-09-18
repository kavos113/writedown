package openapi

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /ping)
func (Server) GetPing(ctx echo.Context) error {
	m := "pong"
	resp := Pong{
		Message: &m,
	}

	return ctx.JSON(http.StatusOK, resp)
}
