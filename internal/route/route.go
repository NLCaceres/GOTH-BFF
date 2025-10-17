package route

import (
	"context"
	"github.com/NLCaceres/goth-example/internal/handler"
	"github.com/NLCaceres/goth-example/internal/handler/queryapi"
	"github.com/NLCaceres/goth-example/internal/model"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/NLCaceres/goth-example/internal/view/index"
	"github.com/NLCaceres/goth-example/internal/view/items"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

// NOTE: Public funcs in Go start with a capital 1st letter, no keyword needed

func Routes(app *echo.Echo) {
	app.GET("/", handler.RenderView)
	app.GET("/:name", func(c echo.Context) error {
		name := c.Param("name")
		newCtx := context.WithValue(c.Request().Context(), "param", name)
		c.SetRequest(c.Request().WithContext(newCtx))
		vm := index.ViewModel{Title: name, CssPaths: []string{"css/item_list.css"}}
		return handler.RenderHTMLView(c, items.ListPage(model.ManyMockItems()), vm)
	})

	apiRoutes := strings.Split(os.Getenv("APP_ROUTES"), ",") // Get comma-delim'd route paths
	routeMap := stringy.Map(os.Getenv("ROUTE_MAP"))
	for _, route := range apiRoutes {
		routePath := "/" + route
		routeFormattedPath := "/" + stringy.PresenterMapValue(routeMap, route)

		app.GET(routePath, queryapi.Call)
		app.GET(routeFormattedPath, queryapi.Call)
	}
}
