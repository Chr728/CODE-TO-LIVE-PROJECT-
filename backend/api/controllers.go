package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//SearchPlaylist
//fetch Spotify Playlist
//Inputs playlistsID string
// Return Playlist Details
func (s *AuthorizationResponse) SearchPlaylist() interface{} {
	auth := fmt.Sprintf("Bearer %s", s.AccessToken)
	var params string
	params = "3cEYpjA9oz9GiPac4AsH4n"
	data := "?market=ES&fields=description%2C%20external_urls%2C%20id%2C%20images%2C%20name%2C%20owner%2C%20"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(body,&v)
	if err != nil {
		log.Fatal(err.Error())
	}

	return v

}

