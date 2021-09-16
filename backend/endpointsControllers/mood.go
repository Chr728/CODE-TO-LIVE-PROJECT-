package endpointsControllers

import (
	"FitMood/backend/api"
	"github.com/labstack/echo"
	"log"
)

type ApiInterface struct {
	spotify               	api.InterfaceSpotify
}


func (s *ApiInterface) SpotifyAPI(clientID, clientSecret string)  {
	spotifyApi := api.New()
	spotifyInterface := api.InterfaceSpotify(spotifyApi)
	err := spotifyInterface.Authorization(clientID, clientSecret)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.spotify = spotifyInterface

}

func (s *ApiInterface) GetMoodPlaylist(c echo.Context) error {
	playlist := s.spotify.SearchPlaylist()

	return c.JSONPretty(200, playlist, "" )
}