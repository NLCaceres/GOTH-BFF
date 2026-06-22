package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/htmx"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/labstack/echo/v5"
	"strings"
)

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
	return HtmxPayload(c, htmx.Data(index.HTML(items.ListPage(itemsVm), vm)), listPage)
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
