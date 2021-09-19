package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PlaylistModel struct {
	Description  string `json:"description"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Id     string `json:"id"`
	Name string `json:"name"`
	Event	string	`json:"event"`
	Images []struct {
		Height interface{} `json:"height"`
		Url    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	ImageUrl	string	`json:"image_url"`
}





//SearchPlaylist
//fetch Spotify Playlist
//Inputs playlistsID string
// Return Playlist Details
func (s *AuthorizationResponse) SearchPlaylist(params string) PlaylistModel {
	auth := fmt.Sprintf("Bearer %s", s.AccessToken)

	data := "?market=ES&fields=description%2C%20external_urls%2C%20id%2C%20images%2C%20name%2C%20images%2C%20"
	//owner%2C%20"
	endpoint := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s%s", params, data)


	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	//fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	//Close response string
	defer resp.Body.Close()

	//get response string
	var v interface{}
	var playlistModel PlaylistModel
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}


	err = json.Unmarshal(body, &playlistModel)
	err = json.Unmarshal(body,&v)
	if err != nil {
		log.Fatal(err.Error())
	}
	if playlistModel.Images != nil{
		playlistModel.ImageUrl = playlistModel.Images[0].Url
	}

	return playlistModel
}

