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
		Description: "how to flex",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{
			"https://open.spotify.com/playlist/3cEYpjA9oz9GiPac4AsH4n",
		},
		Id:   "3cEYpjA9oz9GiPac4AsH4n",
		Name: "A brand new world",
		Event: "Running",
		ImageUrl: "https://i.scdn.co/image/ab67706c0000bebb8d0ce13d55f634e290f744ba",
	}
	playlistLibData = append(playlistLibData,happyPlaylist)

	//happy 200
	happyPlaylist2 := make(map[string] api.PlaylistModel)
	happyPlaylist2["happy"] =  api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{"https://open.spotify.com/playlist/39oH8fid7H9se9tEP6NdGN"},
		Id:   "39oH8fid7H9se9tEP6NdGN",
		Name: "Drake -> Greatest Hits",
		Event: "Hiking",
		ImageUrl: "https://mosaic.scdn.co/640/ab67616d0000b2733dc98872a3cf00117e4623e5ab67616d0000b27358599e61e8a0487443d31bb6ab67616d0000b2739416ed64daf84936d89e671cab67616d0000b273d6adfbd3f091d6f2b9af6507",
	}

	playlistLibData = append(playlistLibData,happyPlaylist2)
	//sad
	sadPlaylist := make(map[string]api.PlaylistModel)
	sadPlaylist["sad"] = api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{
			"https://open.spotify.com/playlist/0KsHowVe17pYWk8HH2ZqwS\n",
		},
		Id:   "0KsHowVe17pYWk8HH2ZqwS",
		Name: "Action Film (feat. Brymo)",
		Event: "Yoga",
	}
	playlistLibData = append(playlistLibData,sadPlaylist)

	//chilling
	chillingPlaylist := make(map[string]api.PlaylistModel)
	chillingPlaylist["chilling"] = api.PlaylistModel{
		Description: "The tracks heating up the continent right now! Cover: <a href=\"https://open.spotify.com/artist/687cZJR45JO7jhk1LHIbgq?si=yldD7_1gS3WswUoW7Avfhg&dl_branch=1\">Tems </a>",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{
			"https://open.spotify.com/playlist/37i9dQZF1DWYkaDif7Ztbp",
		},
		Id:   "37i9dQZF1DWYkaDif7Ztbp",
		Name: "African Heat",
		Event: "boxing",
		ImageUrl: "https://i.scdn.co/image/ab67706f000000037e62dc01ca5fd2a6860f835b",
	}
	playlistLibData = append(playlistLibData,chillingPlaylist)

	//balling
	ballingPlaylist := make(map[string]api.PlaylistModel)
	ballingPlaylist["balling"] = api.PlaylistModel{
		Description: "",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{
			"https://open.spotify.com/playlist/7gEukSQ9CEsZk6XDHiomvp",
		},
		Id:   "7gEukSQ9CEsZk6XDHiomvp",
		Name: "Rihanna concert vibez",
		Event: "jumping",
	}
	playlistLibData = append(playlistLibData,chillingPlaylist)

	//joyful 200
	joyPlaylist := make(map[string] api.PlaylistModel)
	joyPlaylist["joyful"] =  api.PlaylistModel{
		Description: "Featuring 'Stay With Me', 'Too Good At Goodbyes', 'I'm Not The Only One', 'Dancing With A Stranger', 'Promises' and more. Listen now.",
		ExternalUrls: struct {
			Spotify string `json:"spotify"`
		}{"https://open.spotify.com/playlist/1DXLGZJNUv2yMbTQH4VeLS"},
		Id:   "1DXLGZJNUv2yMbTQH4VeLS",
		Name: "Like I Can - Sam Smith Complete Playlist",
		Event: "Hiking",
		ImageUrl: "https://mosaic.scdn.co/640/ab67616d0000b2733dc98872a3cf00117e4623e5ab67616d0000b27358599e61e8a0487443d31bb6ab67616d0000b2739416ed64daf84936d89e671cab67616d0000b273d6adfbd3f091d6f2b9af6507",
	}
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