package route

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

// NOTE: Public funcs in Go start with a capital 1st letter, no keyword needed
func Routes(app *echo.Echo) {
	routeMap := mapFromString(os.Getenv("ROUTE_MAP"))

	appRoutes := strings.Split(os.Getenv("APP_ROUTES"), ",") // Load comma-separated routes from `.env`
	for _, route := range appRoutes {
		routePath := "/" + route
		var routeFormatted string
		if routeReadable, ok := routeMap[route]; ok { // `ok` is false if no value exists for given key
			routeFormatted = routeReadable // No formatting needed if readable version of route exists in map
		} else {
			routeFormatted = util.TitleCase(route)
		}
		routeFormattedPath := "/" + routeFormatted

		handler := func(c echo.Context) error {
			return c.String(http.StatusOK, fmt.Sprintf("Hello %s", routeFormatted))
		}
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
