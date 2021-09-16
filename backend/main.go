package main

import (
	"FitMood/backend/api"
	"FitMood/backend/route"
	"fmt"
	"log"
)

func main() {
	//router config
	e := route.NewRouter()
	e.Logger.Fatal(e.Start(":8008"))


	spotifyApi := api.New()
	spotifyInterface := api.InterfaceSpotify(spotifyApi)
	err := spotifyInterface.Authorization()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(spotifyInterface.SearchPlaylist())
}
