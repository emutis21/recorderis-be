package models

import "time"

const (
	// Token Durations
	AccessTokenDuration            = 15 * time.Minute
	DefaultRefreshTokenDuration    = 24 * time.Hour
	RememberMeRefreshTokenDuration = 30 * 24 * time.Hour

	// Device Types
	DeviceTypeWeb    = "web"
	DeviceTypeMobile = "mobile"
	DeviceTypeTablet = "tablet"

	// Error Messages
	ErrInvalidCredentials = "invalid credentials"
	ErrUserNotFound       = "user not found"
	ErrUserExists         = "user already exists"
	ErrInvalidToken       = "invalid token"
	ErrExpiredToken       = "token expired"
)
