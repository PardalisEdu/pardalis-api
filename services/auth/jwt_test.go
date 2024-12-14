package auth

import (
	"gitlab.com/pardalis/pardalis-api/configs"
	"testing"
	"time"
)

func TestCreateAndVerifyJWT(t *testing.T) {
	tests := []struct {
		name      string
		userApodo string
		wantErr   bool
	}{
		{
			name:      "Valid token creation",
			userApodo: "testUser",
			wantErr:   false,
		},
		{
			name:      "Empty username",
			userApodo: "",
			wantErr:   true,
		},
	}

	secret := []byte(configs.Envs.JWTSecret)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := CreateJWT(secret, tt.userApodo)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			claims, err := VerifyJWT(token, secret)
			if err != nil {
				t.Errorf("VerifyJWT() error = %v", err)
				return
			}

			if claims["userApodo"] != tt.userApodo {
				t.Errorf("VerifyJWT() userApodo = %v, want %v", claims["userApodo"], tt.userApodo)
			}
		})
	}
}

func TestVerifyJWT_InvalidToken(t *testing.T) {
	tests := []struct {
		name    string
		token   string
		wantErr bool
	}{
		{
			name:    "Empty token",
			token:   "",
			wantErr: true,
		},
		{
			name:    "Malformed token",
			token:   "invalid.token.here",
			wantErr: true,
		},
		{
			name:    "Expired token",
			token:   createExpiredToken(),
			wantErr: true,
		},
	}

	secret := []byte(configs.Envs.JWTSecret)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := VerifyJWT(tt.token, secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyJWT() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func createExpiredToken() string {
	secret := []byte(configs.Envs.JWTSecret)
	token, _ := CreateJWT(secret, "testUser")
	time.Sleep(2 * time.Second)
	return token
}
