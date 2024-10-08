package service

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"time"
	"writedown-server/openapi"
	"writedown-server/repository"
)

type Pages interface {
	PostPages(ctx echo.Context, req openapi.PostPagesJSONRequestBody) (openapi.Page, error)
	DeletePagesPageID(ctx echo.Context, pageID int) error
	GetPagesPageID(ctx echo.Context, pageID int) (openapi.Page, error)
	PatchPagesPageID(ctx echo.Context, pageID int, req openapi.PatchPagesPageIDJSONRequestBody) (openapi.Page, error)
	GetPagesPageIDChild(ctx echo.Context, pageID int) ([]openapi.PageAbstract, error)
}

type pages struct {
	repository.PagesRepository
}

func NewPages() Pages {
	return pages{repository.NewPagesRepository()}
}

func (p pages) PostPages(ctx echo.Context, req openapi.PostPagesJSONRequestBody) (openapi.Page, error) {
	parentPath, err := p.PagesRepository.GetPagePath(req.ParentID)
	if err != nil {
		log.Printf("failed to get parent page path: %v", err)
		return openapi.Page{}, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	path := parentPath + "/" + req.Name
	if req.ParentID == 1 {
		path = "/" + req.Name
	}
	ret, err := p.PagesRepository.CreatePage(repository.Page{
		ParentID:    req.ParentID,
		Name:        req.Name,
		Body:        req.Body,
		Path:        path,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatorName: req.Creator,
	})
	if err != nil {
		log.Printf("failed to create page: %v", err)
		return openapi.Page{}, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := openapi.Page{
		Id:        ret.ID,
		Path:      ret.Path,
		Name:      req.Name,
		Body:      ret.Body,
		CreatedAt: ret.CreatedAt,
		UpdatedAt: ret.UpdatedAt,
		Creator:   ret.CreatorName,
	}
	return res, nil
}

func (p pages) DeletePagesPageID(ctx echo.Context, pageID int) error {
	err := p.PagesRepository.DeletePageByID(pageID)
	if err != nil {
		log.Printf("failed to delete page: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	return nil
}

func (p pages) GetPagesPageID(ctx echo.Context, pageID int) (openapi.Page, error) {
	page, err := p.PagesRepository.GetPageByID(pageID)
	if err != nil {
		log.Printf("failed to get page: %v", err)
		return openapi.Page{}, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := openapi.Page{
		Id:        page.ID,
		Path:      page.Path,
		Name:      strings.Split(page.Path, "/")[len(strings.Split(page.Path, "/"))-1],
		Body:      page.Body,
		CreatedAt: page.CreatedAt,
		UpdatedAt: page.UpdatedAt,
		Creator:   page.CreatorName,
	}
	return res, nil
}

func (p pages) PatchPagesPageID(ctx echo.Context, pageID int, req openapi.PatchPagesPageIDJSONRequestBody) (openapi.Page, error) {
	ret, err := p.PagesRepository.UpdatePageByID(pageID, repository.Page{
		ID:          pageID,
		Name:        req.Name,
		Body:        req.Body,
		UpdatedAt:   time.Now(),
		CreatorName: req.Creator,
	})
	if err != nil {
		log.Printf("failed to update page: %v", err)
		return openapi.Page{}, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := openapi.Page{
		Id:        ret.ID,
		Path:      ret.Path,
		Name:      req.Name,
		Body:      ret.Body,
		CreatedAt: ret.CreatedAt,
		UpdatedAt: ret.UpdatedAt,
		Creator:   ret.CreatorName,
	}
	return res, nil
}

func (p pages) GetPagesPageIDChild(ctx echo.Context, pageID int) ([]openapi.PageAbstract, error) {
	pages, err := p.PagesRepository.GetChildrenPages(pageID)
	if err != nil {
		log.Printf("failed to get children pages: %v", err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := make([]openapi.PageAbstract, len(pages))
	for i, page := range pages {
		res[i] = openapi.PageAbstract{
			Id:    page.ID,
			Path:  page.Path,
			Title: page.Name,
		}
	}

	return res, nil
}
