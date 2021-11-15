package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", Hello)
	e.GET("/names/:id", HelloNames)
	e.Start(":8080")
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func HelloNames(c echo.Context) error {
	name := c.Param("id")
	return c.String(http.StatusOK, "Hello "+name)
}
