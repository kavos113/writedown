package handler

import "github.com/labstack/echo/v4"

func (Server) GetPing(ctx echo.Context) error {
	m := "pong"
	resp := Pong{
		Message
}
