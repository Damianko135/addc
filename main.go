package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files from the public folder
	e.Static("/", "public")
	
	e.Logger.Fatal(e.Start(":8080"))
}
