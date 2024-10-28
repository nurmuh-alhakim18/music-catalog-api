package spotify

type SearchResp struct {
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Total  int     `json:"total"`
	Items  []Track `json:"items"`
}

type RecommendationResp struct {
	Items []Track `json:"items"`
}

type Track struct {
	AlbumType        string   `json:"album_type"`
	AlbumTotalTracks int64    `json:"album_total_tracks"`
	AlbumImages      []string `json:"album_images"`
	AlbumName        string   `json:"album_name"`
	AlbumReleaseDate string   `json:"album_release_date"`
	ArtistsName      []string `json:"artists_name"`
	Explicit         bool     `json:"explicit"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	IsLiked          bool     `json:"is_liked"`
}

type TrackActivityRequest struct {
	TrackID string `json:"track_id"`
	IsLiked bool   `json:"is_liked"`
}
