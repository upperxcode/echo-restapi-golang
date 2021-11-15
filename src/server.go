package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type course struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var courses []string

func main() {
	e := echo.New()

	e.GET("/courses", func(c echo.Context) error {
		courseOne := course{
			Id:   1,
			Name: "Curso Práctico de Go",
		}

		return c.JSON(http.StatusOK, courseOne)
	})

	e.Start(":2000")
}
