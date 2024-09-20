package service

import (
	"github.com/labstack/echo/v4"
	"writedown-server/openapi"
)

type Pages struct{}

func NewPages() Pages {
	return Pages{}
}

func (Pages) PostPages(ctx echo.Context, req openapi.NewPage) (openapi.Page, error) {
	return openapi.Page{}, nil
}

func (Pages) DeletePagesPageID(ctx echo.Context, pageID int64) error {
	return nil
}

func (Pages) GetPagesPageID(ctx echo.Context, pageID int) (openapi.Page, error) {
	return openapi.Page{}, nil
}

func (Pages) PatchPagesPageID(ctx echo.Context, pageID int, req openapi.PatchPage) (openapi.Page, error) {
	return openapi.Page{}, nil
}

func (Pages) GetPagesPageIDChild(ctx echo.Context, pageID int) ([]openapi.PageAbstract, error) {
	return []openapi.PageAbstract{}, nil
}
