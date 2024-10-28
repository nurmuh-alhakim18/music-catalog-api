package middleware

import (
	"context"
	"net/http"

	"github.com/nurmuh-alhakim18/music-catalog-api/config"
	"github.com/nurmuh-alhakim18/music-catalog-api/pkg/auth"
	"github.com/nurmuh-alhakim18/music-catalog-api/pkg/utils"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKeyJWT := config.LoadConfig().SecretKeyJWT
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			utils.ResponseError(w, http.StatusUnauthorized, "unauthorized", err)
			return
		}

		userID, err := auth.ValidateJWT(token, secretKeyJWT)
		if err != nil {
			utils.ResponseError(w, http.StatusUnauthorized, "unauthorized", err)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
