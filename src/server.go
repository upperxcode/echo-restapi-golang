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
		return c.String(http.StatusBadRequest, "The indicated course doesn't exist.")
	})

	e.POST("/courses", func(c echo.Context) error {
		new_course := new(course)
		if err := c.Bind(new_course); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		courses = append(courses, *new_course)

		return c.JSON(http.StatusOK, courses)
	})

	e.PUT("/courses/:id", func(c echo.Context) error {
		updated_course := new(course)
		if err := c.Bind(updated_course); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		for i, course := range courses {
			if strconv.Itoa(course.Id) == c.Param("id") {
				courses = append(courses[:i], courses[i+1:]...)
				courses = append(courses, *updated_course)

				return c.JSON(http.StatusOK, courses)
			}
		}
		return c.String(http.StatusBadRequest, "The indicated course doesn't exist.")

	})

	e.DELETE("/courses/:id", func(c echo.Context) error {
		for i, courseitem := range courses {
			if strconv.Itoa(courseitem.Id) == c.Param("id") {
				courses = append(courses[:i], courses[i+1:]...)
				return c.JSON(http.StatusOK, courses)
			}
		}
		return c.String(http.StatusBadRequest, "The indicated course doesn't exist.")
	})

	e.Start(":2000")
}
