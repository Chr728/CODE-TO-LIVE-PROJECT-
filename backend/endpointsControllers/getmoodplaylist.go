package endpointsControllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)



func Homepage( c echo.Context) error{
	return c.JSONPretty(200, "server up", "")
}


//GetMoodPlaylists
//gets all playlistsassociated with a mood
func (s *ApiInterface) GetMoodPlaylists(c echo.Context) error {
	//mood := c.Param("mood")
	params := "happy"

	playlists := s.GetPlaylistbyMood(params)

	if len(playlists) < 1 {
		return c.JSONPretty(200, "playlist not available, mind some suggestions?", "" )
	}

	return c.JSONPretty(200, playlists, "" )
}


//SelectMoodPlaylists
//selects a single playlist from the list of generated playlistsassociated with a mood
func (s *ApiInterface) SelectMoodPlaylists(c echo.Context) error {
	//mood := c.Param("mood")
	params := "happy"

	playlists := s.GetPlaylistbyMood(params)
	var p string
	if len(playlists) < 1 {
		p = "3cEYpjA9oz9GiPac4AsH4n"
	}

	p = playlists[0].(string)

	playlist := s.spotify.SearchPlaylist(p)


	return c.JSONPretty(200, playlist, "" )
}


//AddPlaylistLibrary
//add playlist library to server storage
func (s *ApiInterface) AddPlaylistLibrary(c echo.Context) error{
	mood := c.Param("mood")
	playlistID := c.Param("playlistID")
	fmt.Println(mood,playlistID)
	s.AddPlaylist(mood, playlistID)

	return c.JSON(http.StatusCreated, s.RandomPlaylist())
}

//SuggestPlaylistLibrary
//suggest a few playlist for client
func (s *ApiInterface) SuggestPlaylistLibrary(c echo.Context) error{

	return c.JSONPretty(http.StatusOK,s.RandomPlaylist() , "")
}


