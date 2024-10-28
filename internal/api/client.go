package api

import (
	"net/http"
)

type Client struct {
	Client              http.Client
	spotifyClientID     string
	spotifyClientSecret string
}

func NewClient(spotifyClientID, spotifyClientSecret string) *Client {
	return &Client{
		Client:              http.Client{},
		spotifyClientID:     spotifyClientID,
		spotifyClientSecret: spotifyClientSecret,
	}
}
