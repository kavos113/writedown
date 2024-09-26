package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"writedown-server/openapi"
	"writedown-server/service"
)

func (Server) PostPages(ctx echo.Context) error {
	req := openapi.PostPagesJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := service.NewPages().PostPages(ctx, req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (Server) DeletePagesPageID(ctx echo.Context, pageID int) error {
	err := service.NewPages().DeletePagesPageID(ctx, pageID)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (Server) GetPagesPageID(ctx echo.Context, pageID int) error {
	res, err := service.NewPages().GetPagesPageID(ctx, pageID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (Server) PatchPagesPageID(ctx echo.Context, pageID int) error {
	req := openapi.PatchPagesPageIDJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := service.NewPages().PatchPagesPageID(ctx, pageID, req)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (Server) GetPagesPageIDChild(ctx echo.Context, pageID int) error {
	res, err := service.NewPages().GetPagesPageIDChild(ctx, pageID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
