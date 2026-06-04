package handler

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/http"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/NLCaceres/goth-example/internal/view/reusable"
	"github.com/labstack/echo/v5"
)

func HtmxPayload(c *echo.Context, name string) error {
	isFullRender, err := isFullRender(http.Header{Header: c.Request().Header})
	if err != nil {
		return RenderHTML(c, reusable.TestElem("Error!"))
	}
	if isFullRender {
		vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
		itemsVm := items.ViewModel{Title: name, Items: model.ManyMockItems()}
		return RenderHTMLIndex(c, items.ListPage(itemsVm), vm)
	} else {
		return RenderHTML(c, reusable.TestElem(name))
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
