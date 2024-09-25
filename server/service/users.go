package service

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"writedown-server/openapi"
	"writedown-server/repository"
)

type Users interface {
	PostUsersSignup(ctx echo.Context, req openapi.PostUsersSignupJSONRequestBody) error
	PostUsersLogin(ctx echo.Context, req openapi.PostUsersLoginJSONRequestBody) error
	GetUsersMe(ctx echo.Context) (openapi.Me, error)
}

type users struct {
	repository.UsersRepository
}

func NewUsers() Users {
	return users{}
}

func (u users) PostUsersSignup(ctx echo.Context, req openapi.PostUsersSignupJSONRequestBody) error {
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

func (u users) PostUsersLogin(ctx echo.Context, req openapi.PostUsersLoginJSONRequestBody) error {
	user, err := u.UsersRepository.GetUser(req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
		} else {
			log.Printf("failed to get user: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
		} else {
			log.Printf("failed to compare password: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}
	}

	sess, err := session.Get("session", ctx)
	if err != nil {
		log.Printf("failed to get session: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	sess.Values["userName"] = user.Username
	err = sess.Save(ctx.Request(), ctx.Response())
	if err != nil {
		log.Printf("failed to save session: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	return nil
}

func (u users) GetUsersMe(ctx echo.Context) (openapi.Me, error) {
	return openapi.Me{}, nil
}
