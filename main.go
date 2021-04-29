package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ListenAddr := ":8080"
	if len(os.Args) >= 1 {
		ListenAddr = os.Args[1]
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.GET("/file/echo.jpeg", func(c echo.Context) error {
		headers := c.Request().Header
		for k, v := range headers {
			fmt.Println(k, v)
		}
		return c.File("echo.jpeg")
	})

	e.Logger.Fatal(e.Start(ListenAddr))
}
