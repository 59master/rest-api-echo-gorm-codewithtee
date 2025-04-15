package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// This is the main entry point for the API server.
	// It will start the server and listen for incoming requests.
	// The server will be configured to use the database and other services.
	// It will also handle routing and middleware.
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
