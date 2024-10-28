package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) GetRecommendations(ctx context.Context, limit int, trackID string) (SpotifyRecommendationsResp, error) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	params.Set("market", "ID")
	params.Set("seed_tracks", trackID)

	apiUrl := "https://api.spotify.com/v1/recommendations"
	fullUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return SpotifyRecommendationsResp{}, err
	}

	token, err := c.GetAccessToken()
	if err != nil {
		return SpotifyRecommendationsResp{}, err
	}

	authHeader := fmt.Sprintf("%s %s", token.TokenType, token.AccessToken)
	req.Header.Set("Authorization", authHeader)

	resp, err := c.Client.Do(req)
	if err != nil {
		return SpotifyRecommendationsResp{}, err
	}

	defer resp.Body.Close()

	var response SpotifyRecommendationsResp
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return SpotifyRecommendationsResp{}, err
	}

	return response, nil
}
