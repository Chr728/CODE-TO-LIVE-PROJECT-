package api

type InterfaceSpotify interface {
	Authorization()	error
	SearchPlaylist() interface{}
}
