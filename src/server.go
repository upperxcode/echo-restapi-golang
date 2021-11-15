package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type course struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type allCourses []course

var courses = allCourses{
	{
		Id:   1,
		Name: "Curso Práctico de Go",
	},
	{
		Id:   2,
		Name: "Curso de Docker",
	},
	{
		Id:   3,
		Name: "Curso de Git",
	},
}

func main() {
	e := echo.New()

	e.GET("/courses", func(c echo.Context) error {
		return c.JSON(http.StatusOK, courses)
	})

	e.GET("/courses/:id", func(c echo.Context) error {
		for _, courseitem := range courses {
			if strconv.Itoa(courseitem.Id) == c.Param("id") {
				return c.JSON(http.StatusOK, courseitem)
			}
		}
		return c.JSON(http.StatusOK, courses)
	})

	e.Start(":2000")
}
