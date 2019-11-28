package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func main()  {
	e := echo.New()
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
