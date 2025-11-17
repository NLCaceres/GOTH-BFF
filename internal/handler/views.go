package handler

import (
	"errors"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/fileread"
	"github.com/NLCaceres/goth-example/internal/util/proxy"
	"github.com/NLCaceres/goth-example/internal/util/slice"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RenderView(c echo.Context) error {
	component := index.HTML(index.Home(), index.ViewModel{Title: "Home", CssPaths: nil})
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func RenderHTMLView(c echo.Context, page templ.Component, vm index.ViewModel) error {
	htmlStr, err := templ.ToGoHTML(c.Request().Context(), index.HTML(page, vm))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	// This string conversion should be instant, unlike converting between []byte/string
	return c.HTML(http.StatusAccepted, string(htmlStr))
}

func RenderQuery(c echo.Context) error {
	name := c.Path()[1:]
	res, err := queryapi.Call(name)
	if err != nil {
		return c.NoContent(queryErrCode(err))
	}
	vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
	itemList := slice.SafeMap(res.Documents(), func(d queryapi.Document) model.Item {
		return model.Item{Name: d.Title, URL: d.URL, Description: d.Title + "\n" + string(d.PostTime)}
	})
	itemsVm := items.ViewModel{Title: name, Items: itemList}
	return RenderHTMLView(c, items.ListPage(itemsVm), vm)
}

func queryErrCode(e error) int {
	if errors.As(e, new(proxy.RequestError)) {
		return http.StatusBadGateway
	} else if errors.As(e, new(fileread.Error)) {
		return http.StatusInternalServerError
	} else if errors.Is(e, queryapi.SearchSetterError) {
		return http.StatusNotImplemented
	} else {
		return http.StatusBadRequest
	}
}
