package api

type InterfaceSpotify interface {
	Authorization(clientId, clientSecret string)	error
	SearchPlaylist(params string) PlaylistModel
}
