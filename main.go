package main

import (
	"github.com/labstack/echo/v4"
	"golang_web_programming/membership"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	membership.InitMembershipRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
