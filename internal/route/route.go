package route

import (
	"github.com/NLCaceres/goth-example/internal/handler"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/NLCaceres/goth-example/internal/view/reusable/htmx"
	"github.com/labstack/echo/v5"
	"os"
	"strings"
)

// NOTE: Public funcs in Go start with a capital 1st letter, no keyword needed

func Routes(app *echo.Echo) {
	app.GET("/", handler.RenderView)
	app.GET("/error", func(c *echo.Context) error {
		cssPaths := map[string]string{"pageStylesheet": ""}
		indexPage := index.HTML(index.Error(), index.ViewModel{Title: "Error", CssPaths: cssPaths})
		return handler.HtmxPayload(c, indexPage, htmx.Data(index.Error()).AddTitle("Error"))
	})
	app.GET("/:name", func(c *echo.Context) error {
		name, err := echo.PathParam[string](c, "name")
		if err != nil {
			return err
		}
		listStyle := "css/item_list.css"
		vm := index.ViewModel{Title: name, CssPaths: map[string]string{"pageStylesheet": listStyle}}
		itemsVm := items.ViewModel{Title: name, Items: model.ManyMockItems()}
		listPage := htmx.Data(items.ListPage(itemsVm)).AddTitle(name).AddStyle(listStyle)
		return handler.HtmxPayload(c, index.HTML(items.ListPage(itemsVm), vm), listPage)
	})
	ApiRoutes(app)
}

func ApiRoutes(app *echo.Echo) {
	routeMap := stringy.Map(os.Getenv("ROUTE_MAP"))
	routeAddedMap := make(map[string]bool)
	for _, route := range strings.Split(os.Getenv("APP_ROUTES"), ",") {
		routePath := "/" + route
		routeFormattedPath := "/" + stringy.PresenterMapValue(routeMap, route)

		if !routeAddedMap[routePath] {
			app.GET(routePath, handler.RenderQuery)
			routeAddedMap[routePath] = true
		}
		if !routeAddedMap[routeFormattedPath] {
			app.GET(routeFormattedPath, handler.RenderQuery)
			routeAddedMap[routeFormattedPath] = true
		}
	}
}
