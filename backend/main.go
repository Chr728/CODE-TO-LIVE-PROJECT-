package main

import "github.com/labstack/echo"

func main() {
	// new Echo instance
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(200, "Homepage", "")
	})

	e.Logger.Fatal(e.Start(":8080"))

}
