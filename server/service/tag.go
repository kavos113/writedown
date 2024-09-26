package service

import (
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
}

func (t tag) DeletePagesPageIDTag(ctx echo.Context, pageID int, tagID int) error {
	//TODO implement me
	panic("implement me")
}

func (t tag) GetPagesPageIDTag(ctx echo.Context, pageID int) ([]openapi.Tag, error) {
	//TODO implement me
	panic("implement me")
}

func (t tag) PostPagesPageIDTag(ctx echo.Context, pageID int, req openapi.PostPagesPageIDTagJSONRequestBody) error {
	//TODO implement me
	panic("implement me")
}
