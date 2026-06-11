package handler

import (
	"fmt"
	myHttp "github.com/NLCaceres/goth-example/internal/util/http"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"net/http"
)

func HtmxPayload(c *echo.Context, fullPage, htmx templ.Component) error {
	isFullRender, err := isFullRender(myHttp.Header{Header: c.Request().Header})
	if err != nil { // Header won't work if initial navigation
		c.Response().Header().Add("Hx-Redirect", "/error")
		return c.NoContent(http.StatusBadRequest)
	}
	if isFullRender {
		return RenderHTML(c, fullPage)
	} else {
		return RenderHTML(c, htmx)
	}
}

func isFullRender(h myHttp.Header) (bool, error) {
	switch {
	case (h.IsNavigationFetch() && h.IsSiteNone()) || !h.HasReferer():
		return true, nil
	case ((h.IsCORSFetch() || h.IsSameOriginFetch()) &&
		h.IsSiteAllSame()) || h.IsRefererAny(myHttp.SafeOrigins...):
		return false, nil
	default:
		return false,
			fmt.Errorf("Issue determining if full render needed: FetchMode = %v  "+
				"FetchSite = %v  Referer = %v", h.FetchMode(), h.FetchSite(), h.Referer())
	}
}
