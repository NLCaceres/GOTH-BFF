package handler

import (
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/NLCaceres/goth-example/internal/view/reusable"
	"github.com/labstack/echo/v5"
)

func HtmxPayload(c *echo.Context, name string) error {
	switch fetchMode := c.Request().Header.Get("Sec-Fetch-Mode"); fetchMode {
	case "navigate":
		vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
		itemsVm := items.ViewModel{Title: name, Items: model.ManyMockItems()}
		return RenderHTMLIndex(c, items.ListPage(itemsVm), vm)
	case "cors", "same-origin":
		return RenderHTML(c, reusable.TestElem(name))
	default:
		return RenderHTML(c, reusable.TestElem("Error!"))
	}
}
