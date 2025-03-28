package handler

import (
	"github.com/NLCaceres/goth-example/internal/view"
	"github.com/labstack/echo/v4"
)

func RenderView(c echo.Context) error {
	component := view.HTMLIndex(view.Home(), "Home")
	return component.Render(c.Request().Context(), c.Response().Writer)
}
