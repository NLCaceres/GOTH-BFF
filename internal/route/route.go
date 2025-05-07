package route

import (
	"github.com/NLCaceres/goth-example/internal/handler"
	"github.com/NLCaceres/goth-example/internal/util/stringy"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

// NOTE: Public funcs in Go start with a capital 1st letter, no keyword needed

func Routes(app *echo.Echo) {
	app.GET("/", handler.RenderView)

	apiRoutes := strings.Split(os.Getenv("APP_ROUTES"), ",") // Get comma-delim'd route paths
	routeMap := stringy.Map(os.Getenv("ROUTE_MAP"))
	for _, route := range apiRoutes {
		routePath := "/" + route
		routeFormattedPath := "/" + stringy.PresenterMapValue(routeMap, route)

		app.GET(routePath, handler.ApiPostRequest)
		app.GET(routeFormattedPath, handler.ApiPostRequest)
	}
}
