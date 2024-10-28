package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type SpotifySearchResp struct {
	Tracks Tracks `json:"tracks"`
}

type SpotifyRecommendationsResp struct {
	Tracks []Track `json:"tracks"`
}

type Tracks struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []Track `json:"items"`
}

type Track struct {
	Album    Album    `json:"album"`
	Artists  []Artist `json:"artists"`
	Explicit bool     `json:"explicit"`
	Href     string   `json:"href"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
}

type Album struct {
	AlbumType   string  `json:"album_type"`
	TotalTracks int64   `json:"total_tracks"`
	Images      []Image `json:"images"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
}

type Artist struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type Image struct {
	URL string `json:"url"`
}

func (c *Client) SearchTrack(query string, limit, offset int) (SpotifySearchResp, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", strconv.Itoa(limit))
	params.Set("offset", strconv.Itoa(offset))

	apiUrl := "https://api.spotify.com/v1/search"
	fullUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return SpotifySearchResp{}, err
	}

	token, err := c.GetAccessToken()
	if err != nil {
		return SpotifySearchResp{}, err
	}

	authHeader := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	req.Header.Set("Authorization", authHeader)

	resp, err := c.Client.Do(req)
	if err != nil {
		return SpotifySearchResp{}, err
	}

	defer resp.Body.Close()

	var response SpotifySearchResp
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return SpotifySearchResp{}, err
	}

	return response, nil
}
