package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "static/index.html")
	e.GET("/post/",post)
	e.GET("/get/", get)

	// render your 404 page
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.String(http.StatusNotFound, "not found page")
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func checkDate(from,to string) (error, time.Time,time.Time) {
	const layout  = "2006-01-02T15:04"
	var err error

	dateFrom, errFrom := time.Parse(layout,from)
	dateTo, errTo := time.Parse(layout,to)
	if errFrom != nil || errTo != nil {
		err = errors.New("Invalid format date")
	}
	return err,dateFrom,dateTo
}

func get(c echo.Context) error {
	sf := c.QueryParam("from")
	st := c.QueryParam("to")
	sCount := c.QueryParam("count")

	errDate, from, to := checkDate(sf,st)
	count, errCount := strconv.Atoi(sCount)
	if errDate != nil || errCount != nil || count <= 0 {
		return c.String(http.StatusBadRequest, "Ошибка ввода")
	} else {
		return c.String(http.StatusOK,  getListUsers(from,to,count)  )
	}
}

func post(c echo.Context) error {
	sf := c.FormValue("from")
	st := c.FormValue("to")
	site := c.FormValue("site")

	errDate,from,to := checkDate(sf,st)
	if errDate != nil {
		return c.String(http.StatusBadRequest, "Ошибка ввода")
	} else {
		return c.String(http.StatusOK, writePost(from,to,site) )
	}


}

