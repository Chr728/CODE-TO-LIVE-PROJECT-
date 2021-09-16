package api

type InterfaceSpotify interface {
	Authorization(clientId, clientSecret string)	error
	SearchPlaylist() interface{}
}
