package handler

import (
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderView(c echo.Context) error {
	component := index.HTMLIndex(index.Home(), "Home", []string{"css/index.css"})
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func RenderHTMLView(c echo.Context, page templ.Component, title string, cssPaths []string) error {
	component := index.HTMLIndex(page, title, cssPaths)
	htmlStr, err := templ.ToGoHTML(c.Request().Context(), component)
	if err != nil {
		return c.NoContent(404)
	}
	// This string conversion should be instant, unlike converting between []byte/string
	return c.HTML(202, string(htmlStr))
}
