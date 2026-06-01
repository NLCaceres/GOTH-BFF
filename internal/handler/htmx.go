package handler

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/NLCaceres/goth-example/internal/view/reusable"
	"github.com/labstack/echo/v5"
	"net/http"
)

func HtmxPayload(c *echo.Context, name string) error {
	switch isFullRender, err := isFullRender(c.Request().Header); {
	case err != nil:
		return RenderHTML(c, reusable.TestElem("Error!"))
	case isFullRender:
		vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
		itemsVm := items.ViewModel{Title: name, Items: model.ManyMockItems()}
		return RenderHTMLIndex(c, items.ListPage(itemsVm), vm)
	case !isFullRender:
		return RenderHTML(c, reusable.TestElem(name))
	default:
		return RenderHTML(c, reusable.TestElem("Error!"))
	}
}

func isFullRender(h http.Header) (bool, error) {
	fetchMode := h.Get("Sec-Fetch-Mode")
	fetchSite := h.Get("Sec-Fetch-Site")
	referer := h.Get("Referer")
	switch {
	case (fetchMode == "navigate" && fetchSite == "none") || referer == "":
		return true, nil
	case ((fetchMode == "cors" || fetchMode == "same-origin") &&
		(fetchSite == "same-origin" || fetchSite == "same-site")) ||
		stringy.HasAnyPrefix(referer, "http://localhost"):
		return false, nil
	default:
		return false,
			fmt.Errorf("Issue determining if full render needed: "+
				"FetchMode = %v  FetchSite = %v  Referer = %v", fetchMode, fetchSite, referer)
	}
}
