package router

import (
	"net/http"

	"github.com/nurmuh-alhakim18/music-catalog-api/internal/handlers"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/middleware"
)

func NewRouter(userHandler *handlers.UserHandler, spotifyHandler *handlers.SpotifyHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/register", userHandler.HandlerRegister)
	mux.HandleFunc("POST /api/login", userHandler.HandlerLogin)

	mux.Handle("GET /api/search", middleware.AuthMiddleware(http.HandlerFunc(spotifyHandler.HandlerSearch)))
	mux.Handle("GET /api/recommendations", middleware.AuthMiddleware(http.HandlerFunc(spotifyHandler.HandlerGetRecommendations)))
	mux.Handle("POST /api/track_activities", middleware.AuthMiddleware(http.HandlerFunc(spotifyHandler.HandlerUpsertTrackActivity)))

	return mux
}
