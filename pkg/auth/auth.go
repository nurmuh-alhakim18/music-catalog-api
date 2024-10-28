package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userID uuid.UUID, secretKeyJWT string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "MusicEase",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
		Subject:   userID.String(),
	})

	key := []byte(secretKeyJWT)
	return token.SignedString(key)
}

func ValidateJWT(tokenString, secretKeyJWT string) (uuid.UUID, error) {
	key := []byte(secretKeyJWT)
	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	if claims.Issuer != "MusicEase" {
		return uuid.Nil, errors.New("invalid issuer")
	}

	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID")
	}

	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authHeaders := headers.Values("Authorization")
	if len(authHeaders) > 1 {
		return "", errors.New("multiple authorizatoin headers found")
	}

	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no auth header found in request")
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) < 2 || authParts[0] != "Bearer" || authParts[1] == "" {
		return "", errors.New("invalid auth header")
	}

	return authParts[1], nil
}
