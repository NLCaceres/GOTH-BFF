package handler

import (
	"github.com/NLCaceres/goth-example/internal/util/htmx"
	"github.com/labstack/echo/v5"
	"net/http"
)

// Takes a full page component or the `Hx-Swap` ready partial component and decides
// based on `Hx-` prefixed headers which to send back as a response.
func HtmxPayload(c *echo.Context, fullPage, partial htmx.PageComponent) error {
	if htmx.NewHeader(c.Request().Header).IsHxRequest() {
		return RenderHtmx(c, partial)
	} else {
		return RenderHtmx(c, fullPage)
	}
}

// Takes a htmx.PageComponent wrapping a templ.Component to convert into an HTML string.
// If the conversion fails, an empty 404 status code is returned.
// If the conversion succeeds, the HTML is returned with status code 200.
func RenderHtmx(c *echo.Context, component htmx.PageComponent) error {
	html, err := component.ToHTML(c.Request().Context())
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.HTML(http.StatusOK, html)
}
