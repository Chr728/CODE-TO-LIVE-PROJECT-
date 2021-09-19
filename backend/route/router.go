package route

import (
	"FitMood/backend/config"
	"FitMood/backend/endpointsControllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


var secrets, err = config.LoadSecrets(false)

//NewRouter
//New instance of server
func NewRouter() *echo.Echo{



	//Endpoints handlers Interface
	spotifyApiController := endpointsControllers.NewControllers()
	spotifyApiController.SpotifyAPI(secrets.SpotifyClientID, secrets.SpotifyClientSecret)

	e:= echo.New()
	// new Echo instance


	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//templates

	//Endpoints
	e.GET("/", endpointsControllers.Homepage)
	e.POST("/playlists", spotifyApiController.GetMoodPlaylists)
	e.GET("/playlist", spotifyApiController.SelectMoodPlaylists)
	e.GET("/addplayist", spotifyApiController.AddPlaylistLibrary)
	e.GET("/suggestplayist", spotifyApiController.SuggestPlaylistLibrary)

	//test
	//e.GET("/friends", friendsEndpoint)
	return e
}




