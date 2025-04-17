package models

import "time"

type TokenClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	Exp    int64  `json:"exp"`
}

type RefreshTokenData struct {
	UserID    string    `json:"user_id"`
	TokenID   string    `json:"token_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
