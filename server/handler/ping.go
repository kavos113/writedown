package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"writedown-server/openapi"
)

func (Server) GetPing(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, openapi.Pong{Message: "pong"})
}
