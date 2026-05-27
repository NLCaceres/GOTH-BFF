package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/slice"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"net/http"
)

func RenderView(c *echo.Context) error {
	component := index.HTML(index.Home(), index.ViewModel{Title: "Home", CssPaths: nil})
	return component.Render(c.Request().Context(), c.Response())
}

func RenderHTML(c *echo.Context, component templ.Component) error {
	htmlStr, err := templ.ToGoHTML(c.Request().Context(), component)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	// This string conversion should be instant, unlike converting between []byte/string
	return c.HTML(http.StatusOK, string(htmlStr))
}

func RenderHTMLIndex(c *echo.Context, page templ.Component, vm index.ViewModel) error {
	return RenderHTML(c, index.HTML(page, vm))
}

func RenderQuery(c *echo.Context) error {
	name := c.Path()[1:]
	res, err := queryapi.Call(name)
	var e queryapi.Error
	if errors.As(err, &e) {
		return c.NoContent(e.Code)
	}
	vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
	itemList := slice.SafeMap(res.Documents(), func(d queryapi.Document) model.Item {
		return model.Item{Name: d.Title, URL: d.URL, Description: d.Title + "\n" + string(d.PostTime)}
	})
	itemsVm := items.ViewModel{Title: name, Items: itemList}
	return RenderHTMLIndex(c, items.ListPage(itemsVm), vm)
}
