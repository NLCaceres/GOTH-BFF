package handler

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/htmx"
	myHttp "github.com/NLCaceres/goth-example/internal/util/http"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"net/http"
)

func HtmxPayload(c *echo.Context, fullPage, partial templ.Component) error {
	isFullRender, err := isFullRender(htmx.NewHeader(c.Request().Header))
	if err != nil { // "Hx-Location" can cause infinite re-renders
		c.Response().Header().Add("Hx-Location", `{"path":"/error", "target":"main"}`)
		return c.NoContent(http.StatusBadRequest)
	}
	if isFullRender {
		return RenderHTML(c, fullPage)
	} else {
		return RenderHTML(c, partial)
	}
}

func isFullRender(h htmx.Header) (bool, error) {
	switch {
	case h.IsHxRequest():
		return false, nil
	case (h.IsNavigationFetch() && h.IsSiteNone()) || !h.HasReferer():
		return true, nil
	case ((h.IsCORSFetch() || h.IsSameOriginFetch()) && h.IsSiteAllSame()) ||
		h.IsRefererAny(myHttp.SafeOrigins...):
		return false, nil
	default:
		return false,
			fmt.Errorf("Issue determining if full render needed: FetchMode = %v  "+
				"FetchSite = %v  Referer = %v", h.FetchMode(), h.FetchSite(), h.Referer())
	}
}
