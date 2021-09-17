package endpointsControllers

import (
	"FitMood/backend/api"
	"log"
)

type ApiInterface struct {
	spotify               	api.InterfaceSpotify
	playlistLib				*[]map[string]interface{}
}

func NewControllers() *ApiInterface {
	playlistLibData := make([]map[string]interface{},0)
	//create default mood playlist
	//happy---
	happyPlaylist := make(map[string]interface{})
	happyPlaylist["happy"] = "3cEYpjA9oz9GiPac4AsH4n"
	playlistLibData = append(playlistLibData,happyPlaylist)
	//happy 200
	happyPlaylist2 := make(map[string]interface{})
	happyPlaylist2["happy"] = "playlistHappy"
	playlistLibData = append(playlistLibData,happyPlaylist2)
	//sad
	sadPlaylist := make(map[string]interface{})
	sadPlaylist["sad"] = "playlistSad"
	playlistLibData = append(playlistLibData,sadPlaylist)

	//chilling
	chillingPlaylist := make(map[string]interface{})
	chillingPlaylist["chilling"] = "playlistChilling"
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

func (s *ApiInterface) AddPlaylist(mood, playlist string) {

	//get a playlist
	m := make(map[string]interface{})
	m[mood] = playlist
	// append to list of playist (a slice map of [mood]playlistData)
	*s.playlistLib = append(*s.playlistLib, m)

}

func (s *ApiInterface) GetPlaylistbyMood(mood string) []interface{}{
	//range over backend playlistsdata, and provide most suitable results
	playlists := make([]interface{},0)
	for _,v := range *s.playlistLib{

		if playlist, ok := v[mood]; ok{
			playlists = append(playlists, playlist)
		}
	}

	return playlists
}

func (s *ApiInterface) RandomPlaylist() []interface{}{
	playlists := make([]interface{},0)
	for _, v := range *s.playlistLib{
		playlists = append(playlists, v)
	}
	return playlists
}