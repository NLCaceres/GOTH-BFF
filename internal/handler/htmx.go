package handler

import (
	"github.com/NLCaceres/goth-example/internal/util/htmx"
	"github.com/labstack/echo/v5"
)

func HtmxPayload(c *echo.Context, fullPage, partial htmx.PageComponent) error {
	if htmx.NewHeader(c.Request().Header).IsHxRequest() {
		return RenderHTML(c, partial)
	} else {
		return RenderHTML(c, fullPage)
	}
}
