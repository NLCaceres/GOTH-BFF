package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RenderView(c echo.Context) error {
	component := index.HTML(index.Home(), index.ViewModel{Title: "Home", CssPaths: nil})
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func RenderHTMLView(c echo.Context, page templ.Component, vm index.ViewModel) error {
	htmlStr, err := templ.ToGoHTML(c.Request().Context(), index.HTML(page, vm))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	// This string conversion should be instant, unlike converting between []byte/string
	return c.HTML(http.StatusAccepted, string(htmlStr))
}

func InlineQueries(c echo.Context) error {
	res, err := queryapi.Call(c)
	if err != nil {
		return c.NoContent(queryErrCode(err))
	}
	return c.JSON(http.StatusOK, res)
}

func queryErrCode(e error) int {
	if errors.As(e, new(fileread.FileReadError)) {
		return http.StatusInternalServerError
	} else if errors.Is(e, queryapi.SearchSetterError) {
		return http.StatusNotImplemented
	} else {
		return http.StatusBadRequest
	}
}
