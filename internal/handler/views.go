package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/htmx"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/url"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/labstack/echo/v5"
	"log"
	"net/http"
	"strings"
)

func RenderQuery(c *echo.Context) error {
	res, err := queryapi.Call(url.New(*c.Request().URL))
	var e queryapi.Error
	if errors.As(err, &e) {
		if e.Code == http.StatusUnprocessableEntity {
			return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Path)
		} else {
			return c.NoContent(e.Code)
		}
	}
	itemList := toItems(res.Documents(), strings.Split(c.QueryParam("exclude"), ","))
	listStyle := "/css/item_list.css"
	name := c.Path()[1:]
	vm := index.ViewModel{Title: name, CssPaths: map[string]string{"pageStylesheet": listStyle}}
	page, err := url.URL{URL: *c.Request().URL}.QueryInt("page")
	if err != nil {
		log.Printf("Page %v converted to int %d failed: %v", c.QueryParam("page"), page, err)
		return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Path)
	}
	itemsVm := items.ViewModel{
		Title: name, Items: itemList, CurrentPage: page, PageTotal: 5,
	}
	listPage := htmx.Data(items.ListPage(itemsVm)).AddTitle(name).AddStyle(listStyle)
	return htmx.Response(c, htmx.Data(index.HTML(items.ListPage(itemsVm), vm)), listPage)
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
