package spotify

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/api"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/models/spotify"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/repositories"
)

type SpotifyService struct {
	apiClient *api.Client
	queries   *repositories.Queries
}

func NewSpotifyService(spotifyClientID, spotifyClientSecret string, queries *repositories.Queries) *SpotifyService {
	return &SpotifyService{
		apiClient: api.NewClient(spotifyClientID, spotifyClientSecret),
		queries:   queries,
	}
}

func (s *SpotifyService) SearchTrack(ctx context.Context, query string, pageSize, pageIndex int, userID uuid.UUID) (spotify.SearchResp, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.apiClient.SearchTrack(query, limit, offset)
	if err != nil {
		return spotify.SearchResp{}, err
	}

	trackIDs := []string{}
	for _, item := range trackDetails.Tracks.Items {
		trackIDs = append(trackIDs, item.ID)
	}

	activities, err := s.queries.GetTrackActivitiesForTracks(ctx, repositories.GetTrackActivitiesForTracksParams{
		UserID:   userID,
		TrackIDs: trackIDs,
	})
	if err != nil {
		return spotify.SearchResp{}, err
	}

	activityMap := make(map[string]repositories.TrackActivity)
	for _, activity := range activities {
		activityMap[activity.TrackID] = activity
	}

	items := []spotify.Track{}
	for _, item := range trackDetails.Tracks.Items {
		artistName := []string{}
		for _, artist := range item.Artists {
			artistName = append(artistName, artist.Name)
		}

		images := []string{}
		for _, img := range item.Album.Images {
			images = append(images, img.URL)
		}

		var isLiked bool
		if activity, ok := activityMap[item.ID]; ok {
			isLiked = activity.IsLiked
		} else {
			isLiked = false
		}

		items = append(items, spotify.Track{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImages:      images,
			AlbumName:        item.Album.Name,
			AlbumReleaseDate: item.Album.ReleaseDate,
			ArtistsName:      artistName,
			Explicit:         item.Explicit,
			ID:               item.ID,
			Name:             item.Name,
			IsLiked:          isLiked,
		})
	}

	return spotify.SearchResp{
		Limit:  trackDetails.Tracks.Limit,
		Offset: trackDetails.Tracks.Offset,
		Total:  trackDetails.Tracks.Total,
		Items:  items,
	}, nil
}

func (s *SpotifyService) UpsertTrackActivity(ctx context.Context, userID uuid.UUID, req spotify.TrackActivityRequest) error {
	_, err := s.queries.GetTrackActivity(ctx, repositories.GetTrackActivityParams{
		UserID:  userID,
		TrackID: req.TrackID,
	})
	if err == sql.ErrNoRows {
		err = s.queries.CreateTrackActivities(ctx, repositories.CreateTrackActivitiesParams{
			UserID:  userID,
			TrackID: req.TrackID,
			IsLiked: req.IsLiked,
		})
		if err != nil {
			return errors.New("failed to create track activity")
		}
	} else if err != nil {
		return errors.New("failed to retrieve track activity")
	}

	err = s.queries.UpdateTrackActivities(ctx, repositories.UpdateTrackActivitiesParams{
		UserID:  userID,
		TrackID: req.TrackID,
		IsLiked: req.IsLiked,
	})
	if err != nil {
		return errors.New("failed to update track activity")
	}

	return nil
}

func (s *SpotifyService) GetRecommendations(ctx context.Context, limit int, trackID string, userID uuid.UUID) (spotify.RecommendationResp, error) {
	trackDetails, err := s.apiClient.GetRecommendations(ctx, limit, trackID)
	if err != nil {
		return spotify.RecommendationResp{}, err
	}

	trackIDs := []string{}
	for _, item := range trackDetails.Tracks {
		trackIDs = append(trackIDs, item.ID)
	}

	activities, err := s.queries.GetTrackActivitiesForTracks(ctx, repositories.GetTrackActivitiesForTracksParams{
		UserID:   userID,
		TrackIDs: trackIDs,
	})
	if err != nil {
		return spotify.RecommendationResp{}, err
	}

	activityMap := make(map[string]repositories.TrackActivity)
	for _, activity := range activities {
		activityMap[activity.TrackID] = activity
	}

	items := []spotify.Track{}
	for _, item := range trackDetails.Tracks {
		artistName := []string{}
		for _, artist := range item.Artists {
			artistName = append(artistName, artist.Name)
		}

		images := []string{}
		for _, img := range item.Album.Images {
			images = append(images, img.URL)
		}

		var isLiked bool
		if activity, ok := activityMap[item.ID]; ok {
			isLiked = activity.IsLiked
		} else {
			isLiked = false
		}

		items = append(items, spotify.Track{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImages:      images,
			AlbumName:        item.Album.Name,
			AlbumReleaseDate: item.Album.ReleaseDate,
			ArtistsName:      artistName,
			Explicit:         item.Explicit,
			ID:               item.ID,
			Name:             item.Name,
			IsLiked:          isLiked,
		})
	}

	return spotify.RecommendationResp{
		Items: items,
	}, nil
}
