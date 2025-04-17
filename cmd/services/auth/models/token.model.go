package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string    `json:"user_id"`
	Role   string    `json:"role"`
	Type   TokenType `json:"token_type"`
}

type TokenConfig struct {
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	SigningKey           []byte
	Issuer               string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type RefreshTokenDetails struct {
	TokenID      string    `json:"token_id"`
	UserID       string    `json:"user_id"`
	TokenHash    string    `json:"token_hash"`
	DeviceType   string    `json:"device_type"`
	DeviceName   string    `json:"device_name"`
	IsRememberMe bool      `json:"is_remember_me"`
	ExpiresAt    time.Time `json:"expires_at"`
	LastUsedAt   time.Time `json:"last_used_at"`
}

type SessionInfo struct {
	SessionID      string    `json:"session_id"`
	UserID         string    `json:"user_id"`
	RefreshTokenID string    `json:"refresh_token_id"`
	IPAddress      string    `json:"ip_address"`
	UserAgent      string    `json:"user_agent"`
	LastActivity   time.Time `json:"last_activity"`
}
