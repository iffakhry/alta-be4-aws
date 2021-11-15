package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}
	e := echo.New()
	e.GET("/", Hello)
	e.GET("/names/:id", HelloNames)
	e.GET("/users/:id", HelloUsers)
	e.Start(port)
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func HelloNames(c echo.Context) error {
	name := c.Param("id")
	return c.String(http.StatusOK, "Hello "+name)
}

func HelloUsers(c echo.Context) error {
	name := c.Param("id")
	return c.String(http.StatusOK, "Welcome. Hello users "+name)
}
