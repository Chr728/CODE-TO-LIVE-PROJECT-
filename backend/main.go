package main

import (
	"FitMood/backend/route"
)

func main() {
	// router config
	e := route.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))

}
