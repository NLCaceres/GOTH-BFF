package htmx

import (
	"github.com/labstack/echo/v5"
	"net/http"
)

// Takes a full page component or the `Hx-Swap` ready partial component and decides
// based on `Hx-` prefixed headers which to send back as a response.
func Response(c *echo.Context, fullPage, partial PageComponent) error {
	h := NewHeader(c.Request().Header)
	if h.HxRestoreHistory() {
		return Render(c, fullPage)
	} else if h.IsHxRequest() {
		return Render(c, partial)
	} else {
		return Render(c, fullPage)
	}
}

// Takes a PageComponent wrapping a templ.Component to convert into an HTML string.
// If the conversion fails, an empty 404 status code is returned.
// If the conversion succeeds, the HTML is returned with status code 200.
func Render(c *echo.Context, component PageComponent) error {
	html, err := component.ToHTML(c.Request().Context())
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.HTML(http.StatusOK, html)
}
