package handler

import (
	"context"
	"github.com/NLCaceres/goth-example/internal/view"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderView(c echo.Context) error {
	component := view.HTMLIndex(view.Home(), "Home")
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func RenderHTMLView(c echo.Context, page templ.Component, title string) error {
	component := view.HTMLIndex(page, title)
	newCtx := context.WithValue(c.Request().Context(), "param", c.Param("name"))
	c.SetRequest(c.Request().WithContext(newCtx))
	htmlStr, err := templ.ToGoHTML(c.Request().Context(), component)
	if err != nil {
		return c.NoContent(404)
	}
	// This string conversion should be instant, unlike converting between []byte/string
	return c.HTML(202, string(htmlStr))
}
