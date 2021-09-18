package endpointsControllers


type MoodParams struct {
	Mood	string	`json:"mood" validate:"required"`
	Playlist	string	`json:"playlist" validate:"required"`
}

