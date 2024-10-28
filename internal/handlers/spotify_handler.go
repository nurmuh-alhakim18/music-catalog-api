package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/middleware"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/models/spotify"
	"github.com/nurmuh-alhakim18/music-catalog-api/pkg/utils"
)

type spotifyService interface {
	SearchTrack(ctx context.Context, query string, pageSize, pageIndex int, userID uuid.UUID) (spotify.SearchResp, error)
	UpsertTrackActivity(ctx context.Context, userID uuid.UUID, req spotify.TrackActivityRequest) error
	GetRecommendations(ctx context.Context, limit int, trackID string, userID uuid.UUID) (spotify.RecommendationResp, error)
}

type SpotifyHandler struct {
	spotifyService spotifyService
}

func NewSpotifyHandler(spotifyService spotifyService) *SpotifyHandler {
	return &SpotifyHandler{
		spotifyService: spotifyService,
	}
}

func (h *SpotifyHandler) HandlerSearch(w http.ResponseWriter, r *http.Request) {
	userIDString := fmt.Sprintf("%v", r.Context().Value(middleware.UserIDKey))
	if userIDString == "" {
		utils.ResponseError(w, http.StatusInternalServerError, "user ID not found", nil)
		return
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "failed to parse user ID", err)
		return
	}

	query := r.URL.Query().Get("query")
	pageSizeString := r.URL.Query().Get("page_size")
	pageIndexString := r.URL.Query().Get("page_index")

	if query == "" {
		utils.ResponseError(w, http.StatusBadRequest, "failed to find music", err)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(pageIndexString)
	if err != nil {
		pageIndex = 1
	}

	resp, err := h.spotifyService.SearchTrack(r.Context(), query, pageSize, pageIndex, userID)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "invalid request", err)
		return
	}

	utils.Response(w, http.StatusOK, resp)
}

func (h *SpotifyHandler) HandlerUpsertTrackActivity(w http.ResponseWriter, r *http.Request) {
	userIDString := fmt.Sprintf("%v", r.Context().Value(middleware.UserIDKey))
	if userIDString == "" {
		utils.ResponseError(w, http.StatusInternalServerError, "user ID not found", nil)
		return
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "failed to parse user ID", err)
		return
	}

	var params spotify.TrackActivityRequest
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "invalid input", err)
		return
	}

	err = h.spotifyService.UpsertTrackActivity(r.Context(), userID, params)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SpotifyHandler) HandlerGetRecommendations(w http.ResponseWriter, r *http.Request) {
	userIDString := fmt.Sprintf("%v", r.Context().Value(middleware.UserIDKey))
	if userIDString == "" {
		utils.ResponseError(w, http.StatusInternalServerError, "user ID not found", nil)
		return
	}

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "failed to parse user ID", err)
		return
	}

	trackID := r.URL.Query().Get("track_id")
	limitString := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		limit = 10
	}

	resp, err := h.spotifyService.GetRecommendations(r.Context(), limit, trackID, userID)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "invalid request", err)
		return
	}

	utils.Response(w, http.StatusOK, resp)
}
