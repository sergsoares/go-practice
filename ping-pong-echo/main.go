package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", pong)
	e.Logger.Fatal(e.Start(":8181"))
}

type Pong struct {
	Message string `json:"message"`
}

func pong(c echo.Context) error {
	p := Pong{
		Message: "pong",
	}

	return c.JSON(http.StatusOK, p)
}
