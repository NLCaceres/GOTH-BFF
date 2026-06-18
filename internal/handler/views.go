package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/NLCaceres/goth-example/internal/view/reusable/htmx"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"net/http"
	"strings"
)

func RenderView(c *echo.Context) error {
	cssPaths := map[string]string{"pageStylesheet": ""}
	component := index.HTML(index.Home(), index.ViewModel{Title: "Home", CssPaths: cssPaths})
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
	filters := c.QueryParam("exclude")
	itemList := toItems(res.Documents(), strings.Split(filters, ","))
	listStyle := "css/item_list.css"
	vm := index.ViewModel{Title: name, CssPaths: map[string]string{"pageStylesheet": listStyle}}
	itemsVm := items.ViewModel{Title: name, Items: itemList}
	listPage := htmx.Data(items.ListPage(itemsVm)).AddTitle(name).AddStyle(listStyle)
	return HtmxPayload(c, index.HTML(items.ListPage(itemsVm), vm), listPage)
}

func toItems(docs []queryapi.Document, filters []string) []model.Item {
	items := make([]model.Item, 0, len(docs)) // 0 length, max capacity, backing array small
DocLoop:
	for i, d := range docs {
		for _, filter := range filters {
			if filter != "" && strings.Contains(d.URL, filter) {
				continue DocLoop
			}
		}
		if i > 0 && docs[i].CompanyName == docs[i-1].CompanyName {
			continue DocLoop
		}
		items = append(items, model.Item{
			Name: d.CompanyName, URL: d.URL, Description: d.Title + "\n\n" + string(d.PostTime),
		})
	}
	return items
}
