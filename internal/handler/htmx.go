package handler

import (
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/NLCaceres/goth-example/internal/view/reusable"
	"github.com/labstack/echo/v5"
	"net/http"
)

func HtmxPayload(c *echo.Context, name string) error {
	switch renderMode := getRenderMode(c.Request().Header); renderMode {
	case Full:
		vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
		itemsVm := items.ViewModel{Title: name, Items: model.ManyMockItems()}
		return RenderHTMLIndex(c, items.ListPage(itemsVm), vm)
	case HTMX:
		return RenderHTML(c, reusable.TestElem(name))
	default:
		return RenderHTML(c, reusable.TestElem("Error!"))
	}
}

type RenderMode string

const (
	Full  RenderMode = "full"
	HTMX  RenderMode = "htmx"
	Error RenderMode = "error"
)

func getRenderMode(h http.Header) RenderMode {
	referer := h.Get("Referer")
	fetchMode := h.Get("Sec-Fetch-Mode")
	fetchSite := h.Get("Sec-Fetch-Site")
	switch {
	case (fetchMode == "navigate" &&
		(fetchSite == "same-origin" || fetchSite == "same-site")) || referer == "":
		return Full
	case stringy.HasAnyPrefix(referer, "http://localhost") ||
		((fetchMode == "cors" || fetchMode == "same-origin") &&
			(fetchSite == "same-origin" || fetchSite == "same-site")):
		return HTMX
	default:
		return Error
	}
}
