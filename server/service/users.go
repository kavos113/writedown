package service

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"writedown-server/openapi"
	"writedown-server/repository"
)

type Users struct {
	repository.UsersRepository
}

func NewUsers() Users {
	return Users{}
}

func (u Users) PostUsersSignup(ctx echo.Context, req openapi.PostUsersSignupJSONRequestBody) error {
	count, err := u.UsersRepository.CountUserByUsername(req.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	if count > 0 {
		return echo.NewHTTPError(http.StatusConflict, "username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	err = u.UsersRepository.CreateUser(repository.User{
		Username: req.Username,
		Password: string(hashedPassword),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	return nil
}

func (u Users) PostUsersLogin(ctx echo.Context, req openapi.PostUsersLoginJSONRequestBody) error {
	return nil
}

func (u Users) GetUsersMe(ctx echo.Context) (openapi.Me, error) {
	return openapi.Me{}, nil
}
