package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"writedown-server/openapi"
	"writedown-server/service"
)

func (Server) PostUsersSignup(ctx echo.Context) error {
	req := openapi.PostUsersSignupJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		log.Printf("failed to bind: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if req.Username == "" || req.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username and password are required")
	}

	err := service.NewUsers().PostUsersSignup(ctx, req)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (Server) PostUsersLogin(ctx echo.Context) error {
	req := openapi.PostUsersLoginJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		log.Printf("failed to bind: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err := service.NewUsers().PostUsersLogin(ctx, req)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (Server) GetUsersMe(ctx echo.Context) error {
	res, err := service.NewUsers().GetUsersMe(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
