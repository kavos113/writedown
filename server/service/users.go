package service

import (
	"github.com/labstack/echo/v4"
	"writedown-server/openapi"
)

type Users struct {
}

func NewUsers() Users {
	return Users{}
}

func (u Users) PostUsersSignup(ctx echo.Context, req openapi.PostUsersSignupJSONRequestBody) error {
	return nil
}

func (u Users) PostUsersLogin(ctx echo.Context, req openapi.PostUsersLoginJSONRequestBody) error {
	return nil
}

func (u Users) GetUsersMe(ctx echo.Context) (openapi.Me, error) {
	return openapi.Me{}, nil
}
