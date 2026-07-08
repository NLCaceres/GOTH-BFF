package route

import (
	"github.com/NLCaceres/goth-example/internal/handler"
	"github.com/NLCaceres/goth-example/internal/handler/htmx"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/labstack/echo/v5"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// NOTE: Public funcs in Go start with a capital 1st letter, no keyword needed

func Routes(app *echo.Echo) {
	app.GET("/", func(c *echo.Context) error {
		cssPaths := map[string]string{"pageStylesheet": ""}
		vm := index.ViewModel{Title: "Home", CssPaths: cssPaths}
		return htmx.Render(c, htmx.Data(index.HTML(index.Home(), vm)))
	})
	app.GET("/error", func(c *echo.Context) error {
		cssPaths := map[string]string{"pageStylesheet": ""}
		indexPage := index.HTML(index.Error(), index.ViewModel{Title: "Error", CssPaths: cssPaths})
		return htmx.Response(c, htmx.Data(indexPage), htmx.Data(index.Error()).AddTitle("Error"))
	})
	app.GET("/:name", func(c *echo.Context) error {
		name, err := echo.PathParam[string](c, "name")
		if err != nil {
			return err
		}
		listStyle := "/css/item_list.css"
		vm := index.ViewModel{Title: name, CssPaths: map[string]string{"pageStylesheet": listStyle}}
		page, err := strconv.Atoi(c.QueryParamOr("page", "1"))
		if err != nil {
			log.Printf("Page %v converted to int %d failed: %v", c.QueryParam("page"), page, err)
			return c.Redirect(http.StatusMovedPermanently, c.Request().URL.Path)
		}
		itemsVm := items.ViewModel{
			Title: name, Items: model.ManyMockItems(), CurrentPage: page, PageTotal: 5,
		}
		listPage := htmx.Data(items.ListPage(itemsVm)).AddTitle(name).AddStyle(listStyle)
		return htmx.Response(c, htmx.Data(index.HTML(items.ListPage(itemsVm), vm)), listPage)
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
