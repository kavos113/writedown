// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// wikiページの作成
	// (POST /pages)
	PostPages(ctx echo.Context) error
	// タグに対応するページの一覧を取得
	// (GET /pages/tag)
	GetPagesTag(ctx echo.Context, params GetPagesTagParams) error
	// wikiページの削除
	// (DELETE /pages/{pageID})
	DeletePagesPageID(ctx echo.Context, pageID int) error
	// wikiページの取得
	// (GET /pages/{pageID})
	GetPagesPageID(ctx echo.Context, pageID int) error
	// wikiページの更新
	// (PATCH /pages/{pageID})
	PatchPagesPageID(ctx echo.Context, pageID int) error
	// wikiページの子ページ取得
	// (GET /pages/{pageID}/child)
	GetPagesPageIDChild(ctx echo.Context, pageID int) error
	// ページのタグを削除
	// (DELETE /pages/{pageID}/tag)
	DeletePagesPageIDTag(ctx echo.Context, pageID int, params DeletePagesPageIDTagParams) error
	// ページのタグ取得
	// (GET /pages/{pageID}/tag)
	GetPagesPageIDTag(ctx echo.Context, pageID int) error
	// ページにタグを追加
	// (POST /pages/{pageID}/tag)
	PostPagesPageIDTag(ctx echo.Context, pageID int) error
	// 疎通確認
	// (GET /ping)
	GetPing(ctx echo.Context) error
	// ログイン
	// (POST /users/login)
	PostUsersLogin(ctx echo.Context) error
	// 自身のユーザー情報取得
	// (GET /users/me)
	GetUsersMe(ctx echo.Context) error
	// ユーザー登録
	// (POST /users/signup)
	PostUsersSignup(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostPages converts echo context to params.
func (w *ServerInterfaceWrapper) PostPages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostPages(ctx)
	return err
}

// GetPagesTag converts echo context to params.
func (w *ServerInterfaceWrapper) GetPagesTag(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPagesTagParams
	// ------------- Required query parameter "tagID" -------------

	err = runtime.BindQueryParameter("form", true, true, "tagID", ctx.QueryParams(), &params.TagID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tagID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPagesTag(ctx, params)
	return err
}

// DeletePagesPageID converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePagesPageID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePagesPageID(ctx, pageID)
	return err
}

// GetPagesPageID converts echo context to params.
func (w *ServerInterfaceWrapper) GetPagesPageID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPagesPageID(ctx, pageID)
	return err
}

// PatchPagesPageID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchPagesPageID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchPagesPageID(ctx, pageID)
	return err
}

// GetPagesPageIDChild converts echo context to params.
func (w *ServerInterfaceWrapper) GetPagesPageIDChild(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPagesPageIDChild(ctx, pageID)
	return err
}

// DeletePagesPageIDTag converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePagesPageIDTag(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeletePagesPageIDTagParams
	// ------------- Required query parameter "tagID" -------------

	err = runtime.BindQueryParameter("form", true, true, "tagID", ctx.QueryParams(), &params.TagID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tagID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePagesPageIDTag(ctx, pageID, params)
	return err
}

// GetPagesPageIDTag converts echo context to params.
func (w *ServerInterfaceWrapper) GetPagesPageIDTag(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPagesPageIDTag(ctx, pageID)
	return err
}

// PostPagesPageIDTag converts echo context to params.
func (w *ServerInterfaceWrapper) PostPagesPageIDTag(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "pageID" -------------
	var pageID int

	err = runtime.BindStyledParameterWithOptions("simple", "pageID", ctx.Param("pageID"), &pageID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pageID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostPagesPageIDTag(ctx, pageID)
	return err
}

// GetPing converts echo context to params.
func (w *ServerInterfaceWrapper) GetPing(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPing(ctx)
	return err
}

// PostUsersLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsersLogin(ctx)
	return err
}

// GetUsersMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersMe(ctx)
	return err
}

// PostUsersSignup converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersSignup(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsersSignup(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/pages", wrapper.PostPages)
	router.GET(baseURL+"/pages/tag", wrapper.GetPagesTag)
	router.DELETE(baseURL+"/pages/:pageID", wrapper.DeletePagesPageID)
	router.GET(baseURL+"/pages/:pageID", wrapper.GetPagesPageID)
	router.PATCH(baseURL+"/pages/:pageID", wrapper.PatchPagesPageID)
	router.GET(baseURL+"/pages/:pageID/child", wrapper.GetPagesPageIDChild)
	router.DELETE(baseURL+"/pages/:pageID/tag", wrapper.DeletePagesPageIDTag)
	router.GET(baseURL+"/pages/:pageID/tag", wrapper.GetPagesPageIDTag)
	router.POST(baseURL+"/pages/:pageID/tag", wrapper.PostPagesPageIDTag)
	router.GET(baseURL+"/ping", wrapper.GetPing)
	router.POST(baseURL+"/users/login", wrapper.PostUsersLogin)
	router.GET(baseURL+"/users/me", wrapper.GetUsersMe)
	router.POST(baseURL+"/users/signup", wrapper.PostUsersSignup)

}
