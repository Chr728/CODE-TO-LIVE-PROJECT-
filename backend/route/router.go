package route

import (
	"FitMood/backend/config"
	"FitMood/backend/endpointsControllers"
	"github.com/labstack/echo"
)




var secrets, err = config.LoadSecrets(false)

//NewRouter
//New instance of server
func NewRouter() *echo.Echo{
	spotifyApiController := &endpointsControllers.ApiInterface{}
	spotifyApiController.SpotifyAPI(secrets.SpotifyClientID, secrets.SpotifyClientSecret)


	// new Echo instance
	e := echo.New()

	//Endpoints
	e.GET("/", Homepage)
	e.GET("/getmoodplayist", spotifyApiController.GetMoodPlaylist)

	return e

}

func Homepage( c echo.Context) error{
		return c.JSONPretty(200, "server up", "")
}

func GetMood( c echo.Context) error{
	return c.JSONPretty(200, "get mood request", "")
}