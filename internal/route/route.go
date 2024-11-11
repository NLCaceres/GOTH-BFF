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
	routeMap := make(map[string]string)
	// NOTE: Golang ISN'T functional so no `map`, `filter`, etc. Just need to use `for`
	for _, routeKeyVal := range strings.Split(os.Getenv("ROUTE_MAP"), ",") {
		// Gets each from ["route:routeReadable", "foo:Foo"]
		splitKeyVal := strings.Split(routeKeyVal, ":") // Then ["route", "routeReadable"]
		routePath, routeReadable := splitKeyVal[0], splitKeyVal[1]
		routeMap[routePath] = routeReadable
	}

	appRoutes := strings.Split(os.Getenv("APP_ROUTES"), ",") // Load comma-separated routes from `.env`
	for _, route := range appRoutes {
		routePath := "/" + route

		app.GET(routePath, func(c echo.Context) error {
			var routeFormatted string
			if routeReadable, ok := routeMap[route]; ok { // `ok` is false if no value exists for given key
				routeFormatted = routeReadable // No formatting needed if readable version of route exists in map
			} else {
				routeFormatted = util.TitleCase(route)
			}

			return c.String(http.StatusOK, fmt.Sprintf("Hello %s", routeFormatted))
		})
	}
}
