package main

import (
	"flag"
	"net/http"
	"github.com/labstack/echo"
)

var addr = flag.String("addr", ":3000", "http service address")

func main() {
	flag.Parse()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(*addr))
}