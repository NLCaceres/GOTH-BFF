package handler

import (
	"github.com/NLCaceres/goth-example/internal/view"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderView(c echo.Context) error {
	component := view.HTMLIndex(view.Home(), "Home")
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func RenderHTMLView(c echo.Context) error {
	component := view.HTMLIndex(view.Home(), "Home")
	htmlStr, err := templ.ToGoHTML(c.Request().Context(), component)
	if err != nil {
		return c.NoContent(404)
	}
	return c.HTML(202, string(htmlStr))

}
