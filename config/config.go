package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	DatabaseURL         string
	SecretKeyJWT        string
	SpotifyClientID     string
	SpotifyClientSecret string
}

func LoadConfig() Config {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	secretKeyJWT := os.Getenv("SECRET_KEY_JWT")
	if dbURL == "" {
		log.Fatal("SECRET_KEY_JWT must be set")
	}

	SpotifyClientID := os.Getenv("SPOTIFY_CLIENT_ID")
	if dbURL == "" {
		log.Fatal("SPOTIFY_CLIENT_ID must be set")
	}

	SpotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if dbURL == "" {
		log.Fatal("SPOTIFY_CLIENT_SECRET must be set")
	}

	return Config{
		Port:                port,
		DatabaseURL:         dbURL,
		SecretKeyJWT:        secretKeyJWT,
		SpotifyClientID:     SpotifyClientID,
		SpotifyClientSecret: SpotifyClientSecret,
	}
}
