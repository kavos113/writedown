package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"writedown-server/openapi"
	"writedown-server/service"
)

func (Server) GetPagesTag(ctx echo.Context, params openapi.GetPagesTagParams) error {
	res, err := service.NewTag().GetPagesTag(ctx, params.TagID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (Server) DeletePagesPageIDTag(ctx echo.Context, pageID int, params openapi.DeletePagesPageIDTagParams) error {
	err := service.NewTag().DeletePagesPageIDTag(ctx, pageID, params.TagID)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (Server) GetPagesPageIDTag(ctx echo.Context, pageID int) error {
	res, err := service.NewTag().GetPagesPageIDTag(ctx, pageID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (Server) PostPagesPageIDTag(ctx echo.Context, pageID int) error {
	req := openapi.PostPagesPageIDTagJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		log.Printf("failed to bind: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err := service.NewTag().PostPagesPageIDTag(ctx, pageID, req)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
