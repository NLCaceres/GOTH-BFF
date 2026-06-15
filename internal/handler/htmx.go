package handler

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/htmx"
	apphttp "github.com/NLCaceres/goth-example/internal/util/http"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"net/http"
)

func HtmxPayload(c *echo.Context, fullPage, partial templ.Component) error {
	isFullRender, err := isFullRender(htmx.NewHeader(c.Request().Header))
	if err != nil {
		htmx.NewHeader(c.Response().Header()).AddLocation("/error", "main")
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
	case h.IsHxRequest() ||
		((h.IsCORSFetch() || h.IsSameOriginFetch()) && h.IsSiteAllSame()) || h.IsRefererAny(apphttp.SafeOrigins...):
		return false, nil
	case (h.IsNavigationFetch() && h.IsSiteNone()) || !h.HasReferer():
		return true, nil
	default:
		return false,
			fmt.Errorf("Issue determining if full render needed: HxRequest = %v  FetchMode = %v"+
				"  FetchSite = %v  Referer = %v", h.HxRequest(), h.FetchMode(), h.FetchSite(), h.Referer())
	}
}
