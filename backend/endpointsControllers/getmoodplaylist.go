package endpointsControllers

import (
	"FitMood/backend/api"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)


var frontend = false
func Homepage( c echo.Context) error{
	return MessageResponse(c, http.StatusOK, "home endpoint")
}


//GetMoodPlaylists
//gets all playlistsassociated with a mood
func (s *ApiInterface) GetMoodPlaylists(c echo.Context) error {
	//mood := c.Param("mood")
	params := new(MoodParams)
	err := c.Bind(params)
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}

	if params.Mood == ""{
		params.Mood = "happy"
	}
	if params.Playlist != "" {
		playlist := s.spotify.SearchPlaylist(params.Playlist)
		playlist.Event = params.Fitevent
		s.AddPlaylist(params.Mood,playlist)
		log.Println(playlist)
		plays := make([]api.PlaylistModel,0)
		plays = append(plays, playlist)
		return DataResponse(c, http.StatusOK, plays)
	}


	playlists := s.GetPlaylistbyMood(params.Mood)

	if len(playlists) < 1 {
		return MessageResponseForward(c,http.StatusOK, "playlist not available, mind some suggestions?")
	}

	//return c.JSONPretty(200, playlists, "" )
	return DataResponse(c, http.StatusOK, playlists)
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

	p = playlists[0].Id

	playlist := s.spotify.SearchPlaylist(p)

	//log.Println(playlist)

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
		params.Playlist = "37i9dQZF1DWYkaDif7Ztbp"
	}

	fmt.Println(params.Mood,params.Playlist)

	playlist := s.spotify.SearchPlaylist(params.Playlist)

	log.Println(playlist)

	s.AddPlaylist(params.Mood,playlist)

	return c.JSON(http.StatusCreated, s.RandomPlaylist())
}

//SuggestPlaylistLibrary
//suggest a few playlist for client
func (s *ApiInterface) SuggestPlaylistLibrary(c echo.Context) error{



	return DataResponse(c, http.StatusOK, s.RandomPlaylist())
	//return c.JSONPretty(http.StatusOK,s.RandomPlaylist() , "")
}


