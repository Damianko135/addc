package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strings"
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func handleDev(c echo.Context) error {
	path := c.Request().URL.Path
	if path == "/" {
		path = "/index.html"
	}

	isProd := strings.ToLower(getEnv("PRODUCTION", "false")) == "true"
	// DEV ONLY: set headers to avoid caching
	if !isProd {
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")
	} else {
		// PROD: You *could* set caching headers for hashed assets here
		c.Response().Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	}

	return c.File("public" + path)
}

func main() {
	// Load .env once
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Determine if in production

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handleDev)

	e.GET("/favicon.ico", func(c echo.Context) error {
		return c.NoContent(200)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
