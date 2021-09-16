package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Secrets struct {
	Port               	string
	Environment        	string
	PulsarUrl          	string
	SpotifyClientID     string
	SpotifyClientSecret	string
}

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

// LoadSecrets loads up Secrets from the .env file once.
// If an env file is present, Secrets will be loaded, else it'll be ignored.
func LoadSecrets( debug bool) (*Secrets, error) {

	//fmt.Print(debug)
	if !debug {
		err := Load()
		if err != nil {
			return nil, err
		}
	}

	_secrets := &Secrets{
		Port: os.Getenv("PORT"),
		SpotifyClientID: os.Getenv("client_id"),
		SpotifyClientSecret: os.Getenv("client_secret"),
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_secrets.Port = port
	return _secrets, nil
}