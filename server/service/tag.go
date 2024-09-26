package service

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"writedown-server/openapi"
	"writedown-server/repository"
)

type Tag interface {
	GetPagesTag(ctx echo.Context, tagID int) ([]openapi.PageAbstract, error)
	DeletePagesPageIDTag(ctx echo.Context, pageID int, tagID int) error
	GetPagesPageIDTag(ctx echo.Context, pageID int) ([]openapi.Tag, error)
	PostPagesPageIDTag(ctx echo.Context, pageID int, req openapi.PostPagesPageIDTagJSONRequestBody) error
}

type tag struct {
	repository.TagsRepository
}

func NewTag() Tag {
	return tag{}
}

func (t tag) GetPagesTag(ctx echo.Context, tagID int) ([]openapi.PageAbstract, error) {
	pages, err := t.TagsRepository.GetPagesByTagID(tagID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := make([]openapi.PageAbstract, 0, len(pages))
	for _, page := range pages {
		res = append(res, openapi.PageAbstract{
			Id:    page.ID,
			Path:  page.Path,
			Title: page.Name,
		})
	}

	return res, nil
}

func (t tag) DeletePagesPageIDTag(ctx echo.Context, pageID int, tagID int) error {
	err := t.TagsRepository.DeleteTagByTagID(tagID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return nil
}

func (t tag) GetPagesPageIDTag(ctx echo.Context, pageID int) ([]openapi.Tag, error) {
	tags, err := t.TagsRepository.GetTagsByPageID(pageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, echo.NewHTTPError(http.StatusNotFound, "tag not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	res := make([]openapi.Tag, 0, len(tags))
	for _, tag := range tags {
		res = append(res, openapi.Tag{
			Id:   tag.ID,
			Name: tag.Name,
		})
	}

	return res, nil
}

func (t tag) PostPagesPageIDTag(ctx echo.Context, pageID int, req openapi.PostPagesPageIDTagJSONRequestBody) error {
	_, err := t.TagsRepository.CreateTag(repository.Tag{
		PageID: pageID,
		Name:   req.Name,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return nil
}
