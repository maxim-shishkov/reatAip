package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
	"time"
)

var count int

func main()  {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.Static("/", "static/index.html")
	e.GET("/get/", getListPage )
	e.GET("/post/", getPostPage )

	// render your 404 page
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.String(http.StatusNotFound, "not found page")
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func getListData(from,to time.Time,requiredSize int) string {

	var users []jsonWrite
	usersChan := make(chan jsonWrite)

	count = 0

	go func() {
		for i := 0; i < requiredSize; i++ {
			usersChan <- readData(from,to)
		}
	}()

	for  {
		users = append(users, <- usersChan )
		if count >= requiredSize {
		break
		}
	}

	jitem, _ := json.Marshal( users )
	return string(jitem)
}


func getListPage(c echo.Context) error {
	sf := c.QueryParam("from")
	st := c.QueryParam("to")
	sCount := c.QueryParam("count")

	errDate, from, to := checkInputDate(sf,st)

	count, errCount := strconv.Atoi(sCount)
	if errDate != nil || errCount != nil || count <= 0 {
		return c.String(http.StatusBadRequest, "Ошибка ввода")
	} else {
		return c.String(http.StatusOK,	getListData(from,to,count))
	}
}

func getPostPage(c echo.Context) error {
	strFrom := c.FormValue("from")
	strTo := c.FormValue("to")
	site := c.FormValue("site")

	errDate,from,to := checkInputDate(strFrom,strTo)
	if errDate != nil {
		return c.String(http.StatusBadRequest, "Ошибка ввода")
	} else {
		return c.String(http.StatusOK, writePost(from,to,site) ) //
	}
}








