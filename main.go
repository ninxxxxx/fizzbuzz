package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/ninxxxxx/fizzbuzz/fizzbuzz"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/fizzbuzz/:n", getFizzBuzz)

	e.Logger.Fatal(e.Start(":1323"))
}

func getFizzBuzz(c echo.Context) error {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		fmt.Println(err)
	}
	result := fizzbuzz.FizzBuzz(n)
	// return c.String(http.StatusOK, result)
	return c.JSON(http.StatusOK, map[string]string{"result": result})
}
