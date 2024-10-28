package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type SpotifyTokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c *Client) GetAccessToken() (SpotifyTokenResp, error) {
	apiUrl := "https://accounts.spotify.com/api/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.spotifyClientID)
	data.Set("client_secret", c.spotifyClientSecret)

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return SpotifyTokenResp{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return SpotifyTokenResp{}, err
	}

	defer resp.Body.Close()

	var response SpotifyTokenResp
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return SpotifyTokenResp{}, err
	}

	return response, nil
}
