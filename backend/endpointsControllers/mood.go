package endpointsControllers

import (
	"FitMood/backend/api"
	"log"
)

type ApiInterface struct {
	spotify               	api.InterfaceSpotify
	playlistLib				*[]map[string]api.PlaylistModel
}

func NewControllers() *ApiInterface {
	playlistLibData := make([]map[string]api.PlaylistModel,0)
	//create default mood playlist
	//happy---
	happyPlaylist := make(map[string]api.PlaylistModel)
	//create playlist
	happyPlaylist["happy"] = api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{},
		Id:   "3cEYpjA9oz9GiPac4AsH4n",
		Name: "",
	}
	playlistLibData = append(playlistLibData,happyPlaylist)

	//happy 200
	happyPlaylist2 := make(map[string] api.PlaylistModel)
	happyPlaylist2["happy"] =  api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{},
		Id:   "happyplaylist",
		Name: "",
	}

	playlistLibData = append(playlistLibData,happyPlaylist2)
	//sad
	sadPlaylist := make(map[string]api.PlaylistModel)
	sadPlaylist["sad"] = api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{},
		Id:   "sadplaylist",
		Name: "",
	}
	playlistLibData = append(playlistLibData,sadPlaylist)

	//chilling
	chillingPlaylist := make(map[string]api.PlaylistModel)
	chillingPlaylist["chilling"] = api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{},
		Id:   "chillingplaylist",
		Name: "",
	}
	playlistLibData = append(playlistLibData,chillingPlaylist)


	return &ApiInterface{
		playlistLib: &playlistLibData,
	}
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

func (s *ApiInterface) AddPlaylist(mood string, playlist api.PlaylistModel ) {

	//get a playlist
	m := make(map[string]api.PlaylistModel)
	m[mood] = playlist
	// append to list of playist (a slice map of [mood]playlistData)
	*s.playlistLib = append(*s.playlistLib, m)

}

func (s *ApiInterface) GetPlaylistbyMood(mood string) []api.PlaylistModel{
	//range over backend playlistsdata, and provide most suitable results
	playlists := make([]api.PlaylistModel,0)
	for _,v := range *s.playlistLib{

		if playlist, ok := v[mood]; ok{
			playlists = append(playlists, playlist)
		}
	}

	return playlists
}

func (s *ApiInterface) RandomPlaylist() []map[string]api.PlaylistModel{
	playlists := make([]map[string]api.PlaylistModel,0)
	for _, v := range *s.playlistLib{
		playlists = append(playlists, v)
	}
	return playlists
}