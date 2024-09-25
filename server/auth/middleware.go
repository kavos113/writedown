package auth

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"slices"
)

type Middleware interface {
	UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
	openapi *openapi3.T
}

func NewMiddleware(openapi *openapi3.T) Middleware {
	return middleware{openapi: openapi}
}

func (m middleware) UserAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	router, err := gorillamux.NewRouter(m.openapi)
	if err != nil {
		panic(err)
	}

	return func(ctx echo.Context) error {
		route, _, err := router.FindRoute(ctx.Request())
		if err != nil {
			log.Printf("failed to find route: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}
		tags := route.PathItem.GetOperation(ctx.Request().Method).Tags
		isNeedAuth := !slices.Contains(tags, "login")

		if !isNeedAuth {
			return next(ctx)
		}

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
