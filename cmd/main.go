package main

import (
	"fmt"
	"github.com/NLCaceres/goth-example/internal/route"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	if dotEnvErr := godotenv.Load(); dotEnvErr != nil {
		log.Fatal("Environment not properly loaded") // "log" is prettier than "fmt" by default
	} // NOTE: "log" AND "fmt" print to the terminal BUT Echo's logger easily hides them

	app := echo.New()
	// `Use` must be used & declared BEFORE starting the app
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true, LogURI: true, LogError: true,
		HandleError: true, // Forward errors to the global handler to decide status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil { // Println provides a simple way of concatening strings with vars with spaces injected between
				fmt.Println("REQUEST URL =", v.URI, "&", "REQUEST Status =", v.Status)
			} else { // Printf provides an old-school Python style of interpolating vars into a string BUT SHOULD end with `\n`
				fmt.Printf("ERROR on REQUEST URL = %v with Status %v & Error = %v\n", v.URI, v.Status, v.Error)
			}
			return nil
		},
	}))

	app.Use(middleware.Static("static"))

	route.Routes(app) // Routes must ALSO be declared before `app.Start` is called

	app.Logger.Debug(app.Start("localhost:3000"))
}
