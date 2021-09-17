package main

import "FitMood/backend/route"

func main() {

	e := route.NewRouter()
	e.Logger.Fatal(e.Start(":8001"))

}
