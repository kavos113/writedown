package handler

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		sess, err := session.Get("session", ctx)
		if err != nil {
			log.Printf("failed to get session: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}

		if sess.Values["userName"] == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		ctx.Set("userName", sess.Values["userName"].(string))
		return next(ctx)
	}
}
