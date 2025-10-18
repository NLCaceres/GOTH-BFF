package handler

import (
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
