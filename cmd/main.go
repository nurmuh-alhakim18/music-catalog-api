package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nurmuh-alhakim18/music-catalog-api/config"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/handlers"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/repositories"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/services/spotify"
	"github.com/nurmuh-alhakim18/music-catalog-api/internal/services/user"
	"github.com/nurmuh-alhakim18/music-catalog-api/router"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed making connection with db: %v", err)
	}

	queries := repositories.New(db)
	userService := user.NewUserService(cfg.SecretKeyJWT, queries)
	spotifyService := spotify.NewSpotifyService(cfg.SpotifyClientID, cfg.SpotifyClientSecret, queries)

	userHandler := handlers.NewUserHandler(userService)
	spotifyHandler := handlers.NewSpotifyHandler(spotifyService)

	mux := router.NewRouter(userHandler, spotifyHandler)

	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	log.Printf("Server running on port: %s\n", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
