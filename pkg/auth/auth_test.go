package auth

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	validToken, _ := GenerateJWT(userID, "secret")
	tests := []struct {
		name        string
		tokenString string
		secretKey   string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Valid Token",
			tokenString: validToken,
			secretKey:   "secret",
			wantUserID:  userID,
			wantErr:     false,
		},
		{
			name:        "Invalid Token",
			tokenString: "invalidtokenstring",
			secretKey:   "secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		{
			name:        "Wrong Secret",
			tokenString: validToken,
			secretKey:   "wrongsecret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tt.tokenString, tt.secretKey)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantUserID, gotUserID)
		})
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer validtoken"},
			},
			wantToken: "validtoken",
			wantErr:   false,
		},
		{
			name: "Invalid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"Invalidbearer token"},
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name:      "Missing Authorizaion Header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Empty Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer "},
			},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Multiple Authorization Headers",
			headers: http.Header{
				"Authorization": []string{"Bearer validtoken", "Bearer anothertoken"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := GetBearerToken(tt.headers)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.wantToken, gotToken)
		})
	}
}
