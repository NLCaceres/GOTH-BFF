package handler

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util/http"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)

func HtmxPayload(c *echo.Context, fullPage, htmx templ.Component) error {
	isFullRender, err := isFullRender(http.Header{Header: c.Request().Header})
	if err != nil {
		cssPaths := map[string]string{"pageStylesheet": ""}
		errorPage := index.HTML(index.Error(), index.ViewModel{Title: "Home", CssPaths: cssPaths})
		return RenderHTML(c, errorPage)
	}
	if isFullRender {
		return RenderHTML(c, fullPage)
	} else {
		return RenderHTML(c, htmx)
	}
}

func isFullRender(h http.Header) (bool, error) {
	switch {
	case (h.IsNavigationFetch() && h.IsSiteNone()) || !h.HasReferer():
		return true, nil
	case ((h.IsCORSFetch() || h.IsSameOriginFetch()) &&
		h.IsSiteAllSame()) || h.IsRefererAny(http.SafeOrigins...):
		return false, nil
	default:
		return false,
			fmt.Errorf("Issue determining if full render needed: FetchMode = %v  "+
				"FetchSite = %v  Referer = %v", h.FetchMode(), h.FetchSite(), h.Referer())
	}
}
