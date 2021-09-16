package route

import "github.com/labstack/echo"

func NewRouter() *echo.Echo{

	// new Echo instance
	e := echo.New()

	//Endpoints
	e.GET("/", Homepage)
	e.GET("/getmood", GetMood)

	return e

}

func Homepage( c echo.Context) error{
		return c.JSONPretty(200, "server up", "")
}

func GetMood( c echo.Context) error{
	return c.JSONPretty(200, "get mood request", "")
}