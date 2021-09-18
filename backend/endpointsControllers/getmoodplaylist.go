package endpointsControllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)


var frontend = false
func Homepage( c echo.Context) error{
	return c.JSONPretty(200, "server up", "")
}


//GetMoodPlaylists
//gets all playlistsassociated with a mood
func (s *ApiInterface) GetMoodPlaylists(c echo.Context) error {
	//mood := c.Param("mood")
	params := new(MoodParams)

	if frontend {
		// bind params
		err := c.Bind(params)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}
	}else {
		params.Mood = "happy"
	}

	playlists := s.GetPlaylistbyMood(params.Mood)

	if len(playlists) < 1 {
		return c.JSONPretty(200, "playlist not available, mind some suggestions?", "" )
	}

	return c.JSONPretty(200, playlists, "" )
}


//SelectMoodPlaylists
//selects a single playlist from the list of generated playlistsassociated with a mood
func (s *ApiInterface) SelectMoodPlaylists(c echo.Context) error {
	//mood := c.Param("mood")
	params := new(MoodParams)

	if frontend {
		// bind params
		err := c.Bind(params)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}
	}else {
		params.Mood = "happy"
	}
	playlists := s.GetPlaylistbyMood(params.Mood)

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
	//mood := c.Param("mood")
	params := new(MoodParams)
	if frontend {
		// bind params
		err := c.Bind(params)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}
	}else {
		params.Mood = "happy"
		params.Playlist = "playlistID"
	}

	fmt.Println(params.Mood,params.Playlist)
	s.AddPlaylist(params.Mood, params.Playlist)

	return c.JSON(http.StatusCreated, s.RandomPlaylist())
}

//SuggestPlaylistLibrary
//suggest a few playlist for client
func (s *ApiInterface) SuggestPlaylistLibrary(c echo.Context) error{

	return c.JSONPretty(http.StatusOK,s.RandomPlaylist() , "")
}


