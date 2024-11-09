package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	// `Use` must be used & declared BEFORE starting the app
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true, LogURI: true, LogError: true,
		HandleError: true, // Forward errors to the global handler to decide status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil { // Println provides a simple way of concatening strings with vars with spaces injected between
				fmt.Println("REQUEST URL =", v.URI, "&", "REQUEST Status =", v.Status)
			} else { // Printf provides an old-school Python style of interpolating vars into a string
				fmt.Printf("ERROR on REQUEST URL = %v with Status %v & Error = %v", v.URI, v.Status, v.Error)
			}
			return nil
		},
	}))
	// app.Use(middleware.Logger()) // Original Echo Logger but not as extensible
	app.Logger.Debug(app.Start("localhost:3000"))
}
