package handler

import (
	"github.com/NLCaceres/goth-example/internal/util/htmx"
	"github.com/labstack/echo/v5"
	"net/http"
)

func HtmxPayload(c *echo.Context, fullPage, partial htmx.PageComponent) error {
	if htmx.NewHeader(c.Request().Header).IsHxRequest() {
		html, err := partial.ToHTML(c.Request().Context())
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.HTML(http.StatusOK, html)
	} else {
		html, err := fullPage.ToHTML(c.Request().Context())
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.HTML(http.StatusOK, html)
	}
}
