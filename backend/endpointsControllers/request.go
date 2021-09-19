package endpointsControllers


type MoodParams struct {
	Mood		string	`json:"mood" validate:"required"`
	Fitevent	string	`json:"fitevent" validate:"required"`
	Playlist	string	`json:"playlist" validate:"required"`
}

