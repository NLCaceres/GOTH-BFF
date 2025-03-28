package route

import (
	"github.com/NLCaceres/goth-example/internal/handler"
	"github.com/NLCaceres/goth-example/internal/util"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

// NOTE: Public funcs in Go start with a capital 1st letter, no keyword needed

func Routes(app *echo.Echo) {
	routeMap := mapFromString(os.Getenv("ROUTE_MAP"))

	app.GET("/", handler.RenderView)

	apiRoutes := strings.Split(os.Getenv("APP_ROUTES"), ",") // Get comma-delim'd route paths
	for _, route := range apiRoutes {
		routePath := "/" + route
		var routeFormatted string
		if routeReadable, ok := routeMap[route]; ok { // `ok` = true if value is in map
			routeFormatted = routeReadable // No formatting needed for existing readable version
		} else {
			routeFormatted = util.TitleCase(route)
		}
		routeFormattedPath := "/" + routeFormatted

		handler := func(c echo.Context) error { return handler.ApiPostRequest(c) }
		app.GET(routePath, handler)
		app.GET(routeFormattedPath, handler)
	}
}

func mapFromString(mapString string) map[string]string {
	newMap := make(map[string]string)
	//NOTE: Golang ISN'T functional so no `map`, `filter`, etc. There's only `for`
	for _, keyValPair := range strings.Split(mapString, ",") {
		splitKeyVal := strings.Split(keyValPair, ":") // [key, value]
		if len(splitKeyVal) > 1 {
			key, value := splitKeyVal[0], splitKeyVal[1]
			newMap[key] = value
		}
	}
	return newMap
}
