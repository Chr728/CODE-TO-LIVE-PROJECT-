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

	//Endpoints handlers Interface
	spotifyApiController := endpointsControllers.NewControllers()
	spotifyApiController.SpotifyAPI(secrets.SpotifyClientID, secrets.SpotifyClientSecret)


	// new Echo instance
	e := echo.New()

	//Endpoints
	e.GET("/", endpointsControllers.Homepage)
	e.GET("/getmoodplayist", spotifyApiController.GetMoodPlaylists)
	e.GET("/selectmoodplayist", spotifyApiController.SelectMoodPlaylists)
	e.GET("/addplayist", spotifyApiController.AddPlaylistLibrary)
	e.GET("/suggestplayist", spotifyApiController.SuggestPlaylistLibrary)

	return e
}



